package main

import (
	"math/rand"
	"miningPoolCli/config"
	"miningPoolCli/utils/api"
	"miningPoolCli/utils/boc"
	"miningPoolCli/utils/files"
	"miningPoolCli/utils/gpuwrk"
	"miningPoolCli/utils/helpers"
	"miningPoolCli/utils/initp"
	"miningPoolCli/utils/logreport"
	"miningPoolCli/utils/mlog"
	"miningPoolCli/utils/server"
	"os/exec"
	"strconv"
	"time"
)

var gpuGoroutines []gpuwrk.GpuGoroutine
var globalTasks []api.Task

func startTask(i int, task api.Task) {
	// gpuGoroutines[i].startTimestamp = time.Now().Unix()

	mineResultFilename := "mined_" + strconv.Itoa(task.Id) + ".boc"
	pathToBoc := config.MinerGetter.MinerDirectory + "/" + mineResultFilename

	cmd := exec.Command(
		"./"+config.MinerGetter.MinerDirectory+"/pow-miner-opencl", "-vv",
		"-g"+strconv.Itoa(gpuGoroutines[i].GpuData.GpuId),
		"-p"+strconv.Itoa(gpuGoroutines[i].GpuData.PlatformId),
		"-F"+strconv.Itoa(config.StaticBeforeMinerSettings.BoostFactor),
		"-t"+strconv.Itoa(config.StaticBeforeMinerSettings.TimeoutT),
		config.StaticBeforeMinerSettings.PoolAddress,
		helpers.ConvertHexData(task.Seed),
		helpers.ConvertHexData(task.Complexity),
		config.StaticBeforeMinerSettings.Iterations,
		task.Giver,
		pathToBoc,
	)

	cmd.Stderr = &gpuGoroutines[i].ProcStderr

	unblockFunc := make(chan struct{}, 1)

	var killedByNotActual bool
	var done bool

	cmd.Start()

	go func() {
		cmd.Wait()
		done = true

		if helpers.StringInSlice(mineResultFilename, files.GetDir(config.MinerGetter.MinerDirectory)) {
			// found
			if !killedByNotActual {
				bocFileInHex, _ := boc.ReadBocFileToHex(pathToBoc)

				bocServerResp, err := api.SendHexBocToServer(bocFileInHex, task.Seed, strconv.Itoa(task.Id))
				if err == nil {
					if bocServerResp.Data == "Found" && bocServerResp.Status == "ok" {
						logreport.ShareFound(gpuGoroutines[i].GpuData.Model, gpuGoroutines[i].GpuData.GpuId, task.Id)
					} else {
						logreport.ShareServerError(task, bocServerResp, gpuGoroutines[i].GpuData.GpuId)
					}
				}
			}
			files.RemovePath(pathToBoc)
		}
		enableTask(i)
		unblockFunc <- struct{}{}
	}()

	go func() {
		for !done {
			if checkTaskAlreadyFound(task.Id) {
				killedByNotActual = true
				if err := cmd.Process.Kill(); err != nil {
					mlog.LogError(err.Error())
				}
				break
			}
			time.Sleep(1 * time.Second)
		}
	}()

	<-unblockFunc
}

func checkTaskAlreadyFound(checkId int) bool {
	for _, task := range globalTasks {
		if task.Id == checkId {
			return false
		}
	}
	return true
}

func syncTasks(firstSync *chan struct{}) {
	var firstSyncIsOk bool
	for {
		if t := api.GetTasks().Tasks; len(t) > 0 {
			globalTasks = t
			if !firstSyncIsOk {
				*firstSync <- struct{}{}
				firstSyncIsOk = true
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func enableTask(gpuGoIndex int) {
	if tLen := len(globalTasks); tLen > 0 {
		go startTask(gpuGoIndex, globalTasks[rand.Intn(len(globalTasks))])
	} else {
		mlog.LogError("can't start task, because the len of globalTasks <= 0")
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	gpus := initp.InitProgram()

	gpuGoroutines = make([]gpuwrk.GpuGoroutine, len(gpus))

	firstSync := make(chan struct{})
	go syncTasks(&firstSync)
	<-firstSync

	for gpuGoIndex := range gpuGoroutines {
		gpuGoroutines[gpuGoIndex].GpuData = gpus[gpuGoIndex]
		enableTask(gpuGoIndex)
	}

	if config.NetSrv.RunThis {
		go server.Entrypoint(&gpuGoroutines)
	}

	for {
		time.Sleep(60 * time.Second)
		gpuwrk.CalcHashrate(&gpuGoroutines)
	}
}
