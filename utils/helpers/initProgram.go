/*
miningPoolCli – open-source tonuniverse mining pool client

Copyright (C) 2021 Alexander Gapak
Copyright (C) 2021 Kirill Glushakov
Copyright (C) 2021 Roman Klimov

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

package helpers

import (
	_ "embed"
	"flag"
	"fmt"
	"miningPoolCli/config"
	"miningPoolCli/utils/miniLogger"
	"os"
	"runtime"
)

func InitProgram() {
	config.Configure()

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, config.Texts.GlobalHelpText)
	}

	flag.StringVar(&config.ServerSettings.AuthKey, "auth-key", "", "")
	flag.StringVar(&config.ServerSettings.MiningPoolServerURL, "url", "https://pool.tonuniverse.com", "")
	flag.Parse()

	switch "" {
	case config.ServerSettings.AuthKey:
		miniLogger.LogFatal("Flag -auth-key is required; for help run with -h flag")
	}

	miniLogger.LogText(config.Texts.Logo)
	miniLogger.LogText(config.Texts.WelcomeAdditionalMsg)

	os, architecture := runtime.GOOS, runtime.GOARCH

	if os == "windows" {
		miniLogger.LogFatal("Unsupported OS detected: " + "Windows")
	} else if os == "darwin" {
		miniLogger.LogFatal("Unsupported OS detected: " + "Mac OS")
	} else if os == "linux" && architecture == "amd64" {
		miniLogger.LogOk("Supported OS detected: " + os + "/" + architecture)
	} else {
		miniLogger.LogFatal("Unsupported OS detected: " + os + "/" + architecture)
	}
	miniLogger.LogInfo("Using mining pool API url: " + config.ServerSettings.MiningPoolServerURL)
	config.OS.OperatingSystem, config.OS.Architecture = os, architecture
}