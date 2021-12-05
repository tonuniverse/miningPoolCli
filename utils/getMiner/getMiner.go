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

package getMiner

import (
	"fmt"
	"io/ioutil"
	"miningPoolCli/config"
	"miningPoolCli/utils/helpers"
	"miningPoolCli/utils/miniLogger"
	"os"

	"github.com/cavaliercoder/grab"
)

func getDir(path string) []string {
	listDir, err := ioutil.ReadDir(path)
	if err != nil {
		miniLogger.LogFatalStackError(err)
	}

	var files []string
	for _, file := range listDir {
		files = append(files, file.Name())
	}

	return files
}

func removePath(strDir string) {
	if err := os.RemoveAll(strDir); err != nil {
		miniLogger.LogFatalStackError(err)
	}
}

func UbubntuGetMiner() {
	miniLogger.LogInfo("Starting to download the miner for a linux system")

	if helpers.StringInSlice(config.MinerGetter.MinerDirectory, getDir(".")) {
		miniLogger.LogInfo("\"" + config.MinerGetter.MinerDirectory + "\"" + " already exists; it will be removed")

		removePath(config.MinerGetter.MinerDirectory)
	}

	getFileResp, err := grab.Get(".", config.MinerGetter.UbuntuMinerRelUrl)
	if err != nil {
		miniLogger.LogFatalStackError(err)
	}

	if helpers.StringInSlice(config.MinerGetter.UbubntuFileName, getDir(".")) {
		miniLogger.LogOk("Download completed \"" + getFileResp.Filename + "\"")
	} else {
		miniLogger.LogFatal("Something went wrong. " + config.MinerGetter.UbubntuFileName + " not found in this catalog")
	}

	if err := os.Mkdir(config.MinerGetter.MinerDirectory, 0755); err != nil {
		miniLogger.LogFatal(err.Error())
	}

	r, err := os.Open(config.MinerGetter.UbubntuFileName)
	if err != nil {
		fmt.Println(err)
	}
	helpers.ExtractTarGz(r, config.MinerGetter.MinerDirectory)

	os.Chmod(config.MinerGetter.MinerDirectory+"/"+"pow-miner-opencl", 0700)

	removePath(config.MinerGetter.UbubntuFileName)

	if helpers.StringInSlice("pow-miner-opencl", getDir(config.MinerGetter.MinerDirectory)) {
		miniLogger.LogOk("The miner is saved in the directory: " + "\"" + config.MinerGetter.MinerDirectory + "\"")
	} else {
		miniLogger.LogFatal("Something went wrong. Miner not found in" + "\"" + config.MinerGetter.MinerDirectory + "\"")
	}
}
