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

package gpuwrk

import (
	"bytes"
	"miningPoolCli/config"
	"miningPoolCli/utils/mlog"
	"os/exec"
	"strconv"
	"strings"
)

type GPUstruct struct {
	GpuId      int    `json:"device_id"`
	Model      string `json:"device_name"`
	PlatformId int    `json:"platform_id"`
}

type GpuGoroutine struct {
	GpuData GPUstruct
	// startTimestamp int64
	CurrentHashrate int

	ProcStderr bytes.Buffer
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

	mlog.LogInfo("Found GPUs:")
	for model, count := range dict {
		mlog.LogInfo("x" + strconv.Itoa(count) + " " + model + "\n")
	}
}

func searchGpusWithRegex() ([]GPUstruct, error) {
	cmd := exec.Command("./" + config.MinerGetter.MinerDirectory + "/pow-miner-opencl")

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Start()
	cmd.Wait()

	matches := config.MRgxKit.FindGPUPat.FindAllString(stderr.String(), -1)

	var gpusArray []GPUstruct

	for _, v := range matches {
		gpuModel := strings.TrimSpace(
			config.MRgxKit.ReplaceEndGPU.ReplaceAllString(config.MRgxKit.ReplaceStartGPU.ReplaceAllString(v, ""), ""),
		)

		panddId := config.MRgxKit.FindIntIds.FindAllString(v, -1)
		// TODO: handle errors
		platformId, _ := strconv.Atoi(strings.Replace(panddId[0], "#", "", -1))
		deviceId, _ := strconv.Atoi(strings.Replace(panddId[1], "#", "", -1))

		gpusArray = append(
			gpusArray,
			GPUstruct{
				GpuId:      deviceId,
				Model:      gpuModel,
				PlatformId: platformId,
			},
		)
	}

	return gpusArray, nil
}
func SearchGpus() []GPUstruct {
	nvidiaGpus, _ := searchGpusWithRegex()
	return nvidiaGpus
}