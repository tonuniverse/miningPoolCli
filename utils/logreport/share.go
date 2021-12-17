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

package logreport

import (
	"fmt"
	"miningPoolCli/utils/api"
	"miningPoolCli/utils/mlog"
	"strconv"
)

func ShareFound(gpuModel string, gpuId int, taskId int) {
	mlog.LogOk(fmt.Sprintf(
		"Share FOUND on \"%s\" | gpu id: %s; task id: %s",
		gpuModel, strconv.Itoa(gpuId), strconv.Itoa(taskId),
	))
}

func ShareServerError(task api.Task, bocResp api.SendHexBocToServerResponse, gpuId int) {
	mlog.LogPass()
	mlog.LogError("Share found but server didn't accept it")
	mlog.LogError("----- Server error response for task with id " + strconv.Itoa(task.Id) + ":")
	mlog.LogError("-Status: " + bocResp.Status)
	mlog.LogError("-Code: " + strconv.Itoa(bocResp.Code))
	mlog.LogError("-Data: " + bocResp.Data)
	mlog.LogError("-Hash: " + bocResp.Hash)
	mlog.LogError("-Complexity: " + bocResp.Complexity)
	mlog.LogError("----- Local data")
	mlog.LogError("-GPU ID: " + strconv.Itoa(gpuId))
	mlog.LogError("-Seed: " + task.Seed)
	mlog.LogError("-Complexity: " + task.Complexity)
	mlog.LogPass()
}
