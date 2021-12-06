/*
miningPoolCli – open-source tonuniverse mining pool client

Copyright (C) 2021 tonuniverse.com

This file is part of miningPoolCli.

miningPoolCli is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

miningPoolCli is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with miningPoolCli.  If not, see <https://www.gnu.org/licenses/>.
*/

package gpuUtils

import (
	"encoding/json"
	"log"
	"miningPoolCli/config"
	"miningPoolCli/utils/miniLogger"
	"os/exec"
	"strconv"
	"strings"
)

// func serachNvidiaGpus() ([]GPUstruct, error) {
// 	// serch nvidia GPUs on linux ONLY

// 	prg := "ls"
// 	nvidiaLinuxPath := "/proc/driver/nvidia/gpus"

// 	cmd := exec.Command(prg, nvidiaLinuxPath)
// 	stdout, err := cmd.Output()
// 	if err != nil {
// 		log.Fatal(err.Error())
// 		return nil, nil
// 	}

// 	gpuFilesList := strings.Fields(string(stdout))
// 	var gpuS []GPUstruct
// 	for i := 0; i < len(gpuFilesList); i++ {
// 		cmd := exec.Command("cat", "/proc/driver/nvidia/gpus/"+gpuFilesList[i]+"/information")
// 		stdout, err = cmd.Output()
// 		if err != nil {
// 			log.Fatal(err.Error())
// 			return nil, nil
// 		}
// 		cmdStr := string(stdout)
// 		gpuModel := strings.TrimSpace(strings.Split(strings.Split(cmdStr, ":")[1], "\n")[0])
// 		gpuId, _ := strconv.Atoi(strings.TrimSpace(strings.Split(strings.Split(cmdStr, ":")[11], "\n")[0]))
// 		gpuS = append(gpuS, GPUstruct{gpuId, gpuModel})
// 	}

// 	return gpuS, nil
// }

type GPUstruct struct {
	GpuId      int    `json:"device_id"`
	Model      string `json:"device_name"`
	PlatformId int    `json:"platform_id"`
}

func LogGpuList(gpus []GPUstruct) {
	var gpuNames []string

	for i := 0; i < len(gpus); i++ {
		gpuNames = append(gpuNames, gpus[i].Model)
	}

	dict := make(map[string]int)
	for _, num := range gpuNames {
		dict[num] = dict[num] + 1
	}

	var text string

	for model, count := range dict {
		text += "x" + strconv.Itoa(count) + " " + model + "\n"
	}

	miniLogger.LogInfo("Found GPUs:")
	miniLogger.LogInfo(text)
}

func serchGpusWithOpenCLAPI() ([]GPUstruct, error) {
	cmd := exec.Command("./"+config.MinerGetter.MinerDirectory+"/pow-miner-opencl", "-L")
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err.Error())
	}

	var gpusArray []GPUstruct

	strsArray := strings.Split(string(stdout), "\n")
	for i := 0; i < len(strsArray); i++ {
		var j GPUstruct
		err := json.Unmarshal([]byte(strsArray[i]), &j)
		if err != nil {
			continue
		}

		// try ignore intel hd graphics
		if strings.Contains(j.Model, "intel") {
			continue
		}

		gpusArray = append(gpusArray, j)
	}

	if len(gpusArray) == 0 {
		miniLogger.LogFatal("No any GPUs found")
	}

	return gpusArray, nil
}
func SearchGpus() []GPUstruct {
	nvidiaGpus, _ := serchGpusWithOpenCLAPI()
	return nvidiaGpus
}
