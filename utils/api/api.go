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

package api

import (
	"encoding/json"
	"miningPoolCli/config"
	"miningPoolCli/utils/miniLogger"
)

type User struct {
	Id      int    `json:"id"`
	Address string `json:"address"`
	Token   string `json:"token"`
	Balance int    `json:"balance"`
}

type AuthResponse struct {
	User        User   `json:"user"`
	PoolAddress string `json:"pool_address"`
	ServerResponse
}

func Auth() {
	jsonData, _ := json.Marshal(map[string]string{"token": config.ServerSettings.AuthKey})
	bodyResp := SendPostJsonReq(
		jsonData,
		config.ServerSettings.MiningPoolServerURL+"/token",
	)

	var serverResp AuthResponse

	err := json.Unmarshal(bodyResp, &serverResp)
	if err != nil {
		miniLogger.LogFatal(err.Error())
	}

	if serverResp.User.Id != 0 {
		miniLogger.LogOk("Authorization successful\n")
		if serverResp.User.Address != "" {
			miniLogger.LogInfo("Your TON wallet:")
			miniLogger.LogInfo(serverResp.User.Address)
		} else {
			miniLogger.LogInfo("You can set your TON wallet in https://t.me/tonuniversebot")
		}

		config.StaticBeforeMinerSettings.PoolAddress = serverResp.PoolAddress
	} else {
		miniLogger.LogFatal("Auth failed; invalid token")
	}
	miniLogger.LogPass()
}

type Task struct {
	Id            int    `json:"id"`
	Seed          string `json:"seed"`
	NewComplexity string `json:"new_complexity"`
	Giver         string `json:"address"`
}

type GetTasksResponse struct {
	Tasks []Task `json:"tasks"`
	Found int    `json:"found"`
	ServerResponse
}

func GetTasks() GetTasksResponse {
	jsonData, _ := json.Marshal(map[string]string{})
	bodyResp := SendPostJsonReq(
		jsonData,
		config.ServerSettings.MiningPoolServerURL+"/get",
	)

	var results GetTasksResponse

	if err := json.Unmarshal(bodyResp, &results); err != nil {
		miniLogger.LogFatal(err.Error())
	}

	return results
}

type SendHexBocToServerResponse struct {
	ServerResponse
	Hash       string `json:"hash"`
	Complexity string `json:"complexity"`
}

func SendHexBocToServer(hexData string, seed string) SendHexBocToServerResponse {
	jsonData, _ := json.Marshal(map[string]string{
		"hexData":    hexData,
		"dataSource": "minerClient",
		"token":      config.ServerSettings.AuthKey,
		"speed":      "1",
		"seed":       seed,
	})

	bodyResp := SendPostJsonReq(
		jsonData,
		config.ServerSettings.MiningPoolServerURL+"/boc",
	)

	var results SendHexBocToServerResponse
	if err := json.Unmarshal(bodyResp, &results); err != nil {
		miniLogger.LogFatalStackError(err)
	}

	return results
}
