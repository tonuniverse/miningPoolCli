package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"miningPoolCli/config"
	"miningPoolCli/utils/api"
	"miningPoolCli/utils/bocUtils"
	"miningPoolCli/utils/files"
	"miningPoolCli/utils/gpuUtils"
	"miningPoolCli/utils/helpers"
	"miningPoolCli/utils/initp"
	"miningPoolCli/utils/miniLogger"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type gpuGoroutine struct {
	gpuData        gpuUtils.GPUstruct
	startTimestamp int64

	procStdout bytes.Buffer
}

var gpuGoroutines []gpuGoroutine
var globalTasks []api.Task

func startTask(gpuGoIndex int, task api.Task) {
	gpuGoroutines[gpuGoIndex].startTimestamp = time.Now().Unix()

	mineResultFilename := "mined_" + strconv.Itoa(gpuGoroutines[gpuGoIndex].gpuData.GpuId) + ".boc"
	pathToBoc := config.MinerGetter.MinerDirectory + "/" + mineResultFilename

	cmd := exec.Command(
		"./"+config.MinerGetter.MinerDirectory+"/pow-miner-opencl", "-vv",
		"-g"+strconv.Itoa(gpuGoroutines[gpuGoIndex].gpuData.GpuId),
		"-p"+strconv.Itoa(config.StaticBeforeMinerSettings.PlatformID),
		"-F"+strconv.Itoa(config.StaticBeforeMinerSettings.BoostFactor), "-t1920",
		config.StaticBeforeMinerSettings.PoolAddress,
		helpers.ConvertHexData(task.Seed),
		helpers.ConvertHexData(task.Complexity),
		config.StaticBeforeMinerSettings.Iterations,
		task.Giver,
		pathToBoc,
	)

	cmd.Stdout = &gpuGoroutines[gpuGoIndex].procStdout

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
				bocFileInHex, _ := bocUtils.ReadBocFileToHex(pathToBoc)

				bocServerResp := api.SendHexBocToServer(bocFileInHex, task.Seed)
				if bocServerResp.Data == "Found" && bocServerResp.Status == "ok" {
					miniLogger.LogOk(fmt.Sprintf(
						"Share FOUND on \"%s\" | gpu id: %s; task id: %s",
						gpuGoroutines[gpuGoIndex].gpuData.Model,
						strconv.Itoa(gpuGoroutines[gpuGoIndex].gpuData.GpuId),
						strconv.Itoa(task.Id),
					))
				} else {
					miniLogger.LogPass()
					miniLogger.LogError("Share found but server didn't accept it")
					miniLogger.LogError("----- Server error response for task with id " + strconv.Itoa(task.Id) + ":")
					miniLogger.LogError("-Status: " + bocServerResp.Status)
					miniLogger.LogError("-Data: " + bocServerResp.Data)
					miniLogger.LogError("-Hash: " + bocServerResp.Hash)
					miniLogger.LogError("-Complexity: " + bocServerResp.Complexity)
					miniLogger.LogError("----- Local data")
					miniLogger.LogError("-Seed: " + task.Seed)
					miniLogger.LogError("-Complexity: " + task.Complexity)
					miniLogger.LogPass()
				}
			}
			files.RemovePath(pathToBoc)
		}
		enableTask(gpuGoIndex)
		unblockFunc <- struct{}{}
	}()

	go func() {
		for !done {
			if checkTaskAlreadyFound(task.Id) {
				killedByNotActual = true
				if err := cmd.Process.Kill(); err != nil {
					miniLogger.LogError(err.Error())
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
		globalTasks = api.GetTasks().Tasks

		if len(globalTasks) > 0 && !firstSyncIsOk {
			*firstSync <- struct{}{}
			firstSyncIsOk = true
		}

		time.Sleep(3 * time.Second)
	}
}

func enableTask(gpuGoIndex int) {
	go startTask(gpuGoIndex, globalTasks[rand.Intn(len(globalTasks))])
}

func calcHashrate(gpus []gpuGoroutine) {
	var totalHashRate int

	for _, v := range gpus {
		hS := strings.Split(v.procStdout.String(), "\n")

		if len(hS) < 4 {
			return
		}

		sep := strings.Split(hS[len(hS)-2], ".")
		if len(sep) != 2 {
			return
		}

		perHashRate, err := strconv.Atoi(sep[0])
		if err != nil {
			return
		}

		totalHashRate += perHashRate
	}

	miniLogger.LogInfo("Total hashrate: ~" + strconv.Itoa(totalHashRate) + " Mh")
}

func main() {
	rand.Seed(time.Now().Unix())
	gpus := initp.InitProgram()

	gpuGoroutines = make([]gpuGoroutine, len(gpus))

	firstSync := make(chan struct{})
	go syncTasks(&firstSync)
	<-firstSync

	for gpuGoIndex := range gpuGoroutines {
		gpuGoroutines[gpuGoIndex].gpuData = gpus[gpuGoIndex]
		enableTask(gpuGoIndex)
	}

	for {
		calcHashrate(gpuGoroutines)
		time.Sleep(60 * 5 * time.Second)
	}
}
