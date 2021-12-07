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

package files

import (
	"io/ioutil"
	"miningPoolCli/utils/miniLogger"
	"os"
)

func GetDir(path string) []string {
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

func RemovePath(strDir string) {
	if err := os.RemoveAll(strDir); err != nil {
		miniLogger.LogFatalStackError(err)
	}
}