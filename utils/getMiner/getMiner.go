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

package getMiner

import (
	"miningPoolCli/config"
	"miningPoolCli/utils/helpers"
	"miningPoolCli/utils/miniLogger"
	"strings"
)

func splitLs(stdout []byte) []string {
	return strings.Split(string(stdout), "\n")
}

func UbubntuGetMiner() {
	miniLogger.LogInfo("Starting to download the miner for a linux system")

	if helpers.StringInSlice(config.MinerGetter.MinerDirectory, splitLs(helpers.ExecuteSimpleCommand("ls"))) {
		miniLogger.LogInfo("\"" + config.MinerGetter.MinerDirectory + "\"" + " already exists; it will be removed")
	}

	helpers.ExecuteSimpleCommand("rm", "-rf", config.MinerGetter.MinerDirectory)
	helpers.ExecuteSimpleCommand("wget", config.MinerGetter.UbuntuMinerRelUrl)

	if helpers.StringInSlice(config.MinerGetter.UbubntuFileName, splitLs(helpers.ExecuteSimpleCommand("ls"))) {
		miniLogger.LogOk("Download completed \"" + config.MinerGetter.UbubntuFileName + "\"")
	} else {
		miniLogger.LogFatal("Something went wrong. " + config.MinerGetter.UbubntuFileName + " not found in this catalog")
	}

	helpers.ExecuteSimpleCommand("mkdir", config.MinerGetter.MinerDirectory)
	helpers.ExecuteSimpleCommand("tar", "-xvf", config.MinerGetter.UbubntuFileName, "-C", config.MinerGetter.MinerDirectory)
	helpers.ExecuteSimpleCommand("rm", config.MinerGetter.UbubntuFileName)

	if helpers.StringInSlice("pow-miner-opencl", splitLs(helpers.ExecuteSimpleCommand("ls", config.MinerGetter.MinerDirectory))) {
		miniLogger.LogOk("The miner is saved in the directory: " + "\"" + config.MinerGetter.MinerDirectory + "\"")
	} else {
		miniLogger.LogFatal("Something went wrong. Miner not found in" + "\"" + config.MinerGetter.MinerDirectory + "\"")
	}
}
