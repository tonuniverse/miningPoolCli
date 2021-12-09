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

package initp

import (
	"flag"
	"fmt"
	"miningPoolCli/config"
	"miningPoolCli/utils/api"
	"miningPoolCli/utils/getMiner"
	"miningPoolCli/utils/gpuUtils"
	"miningPoolCli/utils/miniLogger"
	"os"
	"runtime"
)

func InitProgram() []gpuUtils.GPUstruct {
	config.Configure()

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, config.Texts.GlobalHelpText)
	}

	flag.StringVar(&config.ServerSettings.AuthKey, "pool-id", "", "")
	flag.StringVar(&config.ServerSettings.MiningPoolServerURL, "url", "https://pool.tonuniverse.com", "")
	flag.BoolVar(&config.UpdateStatsFile, "stats", false, "")
	flag.Parse()

	switch "" {
	case config.ServerSettings.AuthKey:
		miniLogger.LogFatal("Flag -pool-id is required; for help run with -h flag")
	}

	miniLogger.LogText(config.Texts.Logo)
	miniLogger.LogText(config.Texts.WelcomeAdditionalMsg)

	os, architecture := runtime.GOOS, runtime.GOARCH

	if os == config.OSType.Win {
		miniLogger.LogFatal("Unsupported OS detected: " + "Windows")
	} else if os == config.OSType.Macos {
		miniLogger.LogFatal("Unsupported OS detected: " + "Mac OS")
	} else if os == config.OSType.Linux && architecture == "amd64" {
		miniLogger.LogOk("Supported OS detected: " + os + "/" + architecture)
	} else {
		miniLogger.LogFatal("Unsupported OS detected: " + os + "/" + architecture)
	}
	miniLogger.LogInfo("Using mining pool API url: " + config.ServerSettings.MiningPoolServerURL)
	config.OS.OperatingSystem, config.OS.Architecture = os, architecture

	api.Auth()

	getMiner.UbubntuGetMiner()
	gpusArray := gpuUtils.SearchGpus()

	miniLogger.LogPass()
	gpuUtils.LogGpuList(gpusArray)
	miniLogger.LogInfo("Launching the mining processes...")

	return gpusArray
}
