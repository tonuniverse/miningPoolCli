package server

import (
	"fmt"
	"miningPoolCli/utils/gpuwrk"
	"miningPoolCli/utils/mlog"
	"net/http"
	"os"
	"strconv"
	"time"
)

func killHandler(gpuData *[]gpuwrk.GpuGoroutine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" && r.Method != "POST" {
			http.Error(w, string(errJson.MethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		mlog.LogInfo("Received /kill HTTP request")

		for i := 0; i < len(*gpuData); i++ {
			(*gpuData)[i].KeepAlive = false
			pPid := (*gpuData)[i].PPid
			gpuModel := (*gpuData)[i].GpuData.Model
			gpuId := (*gpuData)[i].GpuData.GpuId
			proc, err := os.FindProcess(pPid)
			if err != nil {
				mlog.LogInfo("warning: FindProcess: " + err.Error())
				continue
			}

			if err := proc.Kill(); err != nil {
				mlog.LogInfo("warning: proc.Kill: " + err.Error())
				continue
			}

			mlog.LogOk(fmt.Sprintf(
				"%s (gpuId: %s; pid: %s) - killed",
				gpuModel,
				strconv.Itoa(gpuId),
				strconv.Itoa(pPid),
			))
		}

		defer func() {
			go func() {
				time.Sleep(64 * time.Millisecond)
				mlog.LogOk("Killing main ...")
				os.Exit(1)
			}()
		}()

		w.Write([]byte(`{"status": true}`))
	}
}
