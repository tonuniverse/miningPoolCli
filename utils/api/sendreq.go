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
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"miningPoolCli/config"
	"miningPoolCli/utils/mlog"
	"net/http"
	"strconv"
	"time"
)

type ServerResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
	Code   int    `json:"code"`
}

func SendPostJsonReq(jsonData []byte, serverUrl string) []byte {
	var body []byte = nil
	for attempts := 0; attempts < 5; attempts++ {

		request, _ := http.NewRequest("POST", serverUrl, bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		request.Header.Set("Build-Version", config.BuildVersion)

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Timeout: 2 * time.Second, Transport: tr}

		response, err := client.Do(request)
		if err != nil {
			mlog.LogError(err.Error())
			mlog.LogInfo("Sleep request for 3 sec")
			time.Sleep(3 * time.Second)
			mlog.LogInfo("Attempting to retry the request... [" + strconv.Itoa(attempts+1) + "/" + "3]")
			continue
		}
		defer func() {
			err := response.Body.Close()
			if err != nil {
				mlog.LogFatalStackError(err)
			}
		}()

		body, _ = ioutil.ReadAll(response.Body)
		if attempts > 0 {
			mlog.LogOk("Request sent")
		}
		break
	}
	if body == nil {
		mlog.LogFatal("Attempts to send a request have yielded no results :(")
	}
	return body
}
