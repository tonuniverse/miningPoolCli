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

package getminer

import (
	"miningPoolCli/config"
	"miningPoolCli/utils/files"
	"miningPoolCli/utils/helpers"
	"miningPoolCli/utils/mlog"
	"os"

	"github.com/cavaliercoder/grab"
)

func GetMiner() {
	var minerReleaseURL, minerFileName, executableName string

	switch config.OS.OperatingSystem {
	case config.OSType.Linux:
		minerReleaseURL = config.MinerGetter.UbuntuSettings.ReleaseURL
		minerFileName = config.MinerGetter.UbuntuSettings.FileName
		executableName = config.MinerGetter.UbuntuSettings.ExecutableName
	case config.OSType.Win:
		minerReleaseURL = config.MinerGetter.WinSettings.ReleaseURL
		minerFileName = config.MinerGetter.WinSettings.FileName
		executableName = config.MinerGetter.WinSettings.ExecutableName
	}

	mlog.LogInfo("Starting to download the miner for a " + config.OS.OperatingSystem + " system")

	if helpers.StringInSlice(config.MinerGetter.MinerDirectory, files.GetDir(".")) {
		mlog.LogInfo("\"" + config.MinerGetter.MinerDirectory + "\"" + " already exists; it will be removed")
		files.RemovePath(config.MinerGetter.MinerDirectory)
	}

	getFileResp, err := grab.Get(".", minerReleaseURL)
	if err != nil {
		mlog.LogFatalStackError(err)
	}

	if helpers.StringInSlice(minerFileName, files.GetDir(".")) {
		mlog.LogOk("Download completed \"" + getFileResp.Filename + "\"")
	} else {
		mlog.LogFatal("Something went wrong. " + minerFileName + " not found in this catalog")
	}

	if err := os.Mkdir(config.MinerGetter.MinerDirectory, 0755); err != nil {
		mlog.LogFatal(err.Error())
	}

	switch config.OS.OperatingSystem {
	case config.OSType.Linux:
		r, err := os.Open(minerFileName)
		if err != nil {
			mlog.LogFatal("error while opening '" + minerFileName + "'; " + err.Error())
		}

		files.ExtractTarGz(r, config.MinerGetter.MinerDirectory)
	case config.OSType.Win:
		files.ExtractZip(minerFileName, config.MinerGetter.MinerDirectory)
	}

	if config.OS.OperatingSystem == config.OSType.Linux {
		os.Chmod(config.MinerGetter.MinerDirectory+"/"+executableName, 0700)
	}

	files.RemovePath(minerFileName)

	if helpers.StringInSlice(executableName, files.GetDir(config.MinerGetter.MinerDirectory)) {
		mlog.LogOk("The miner is saved in the directory: " + "\"" + config.MinerGetter.MinerDirectory + "\"")
	} else {
		mlog.LogFatal("Something went wrong. Miner not found in" + "\"" + config.MinerGetter.MinerDirectory + "\"")
	}
}
