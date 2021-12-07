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
	"miningPoolCli/utils/miniLogger"
	"strconv"
)

func ShareFound(gpuModel string, gpuId int, taskId int) {
	miniLogger.LogOk(fmt.Sprintf(
		"Share FOUND on \"%s\" | gpu id: %s; task id: %s",
		gpuModel, strconv.Itoa(gpuId), strconv.Itoa(taskId),
	))
}

func ShareServerError(task api.Task, bocResp api.SendHexBocToServerResponse, gpuId int) {
	miniLogger.LogPass()
	miniLogger.LogError("Share found but server didn't accept it")
	miniLogger.LogError("----- Server error response for task with id " + strconv.Itoa(task.Id) + ":")
	miniLogger.LogError("-Status: " + bocResp.Status)
	miniLogger.LogError("-Code: " + strconv.Itoa(bocResp.Code))
	miniLogger.LogError("-Data: " + bocResp.Data)
	miniLogger.LogError("-Hash: " + bocResp.Hash)
	miniLogger.LogError("-Complexity: " + bocResp.Complexity)
	miniLogger.LogError("----- Local data")
	miniLogger.LogError("-GPU ID: " + strconv.Itoa(gpuId))
	miniLogger.LogError("-Seed: " + task.Seed)
	miniLogger.LogError("-Complexity: " + task.Complexity)
	miniLogger.LogPass()
}
