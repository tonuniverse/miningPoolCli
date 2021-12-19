package server

import (
	"encoding/json"
	"miningPoolCli/config"
	"miningPoolCli/utils/gpuwrk"
	"miningPoolCli/utils/mlog"
	"net/http"
	"time"
)

type info struct {
	Hashrate int `json:"hashrate"`
	gpuwrk.GPUstruct
}

func statHandler(gpuData *[]gpuwrk.GpuGoroutine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" && r.Method != "POST" {
			http.Error(w, string(errJson.MethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		var resp = struct {
			Status        bool   `json:"status"`
			MinerUptime   int64  `json:"miner_uptime"`
			TotalHashrate int    `json:"total_hashrate"`
			Gpus          []info `json:"gpus"`
		}{
			Status:      true,
			MinerUptime: time.Now().Unix() - config.StartProgramTimestamp,
		}

		for i := 0; i < len(*gpuData); i++ {
			g := (*gpuData)[i]
			resp.Gpus = append(resp.Gpus, info{
				GPUstruct: g.GpuData,
				Hashrate:  g.CurrentHashrate,
			})
			resp.TotalHashrate += g.CurrentHashrate
		}

		jsonResp, err := json.Marshal(resp)
		if err != nil {
			mlog.LogFatalStackError(err)
		}
		w.Write(jsonResp)
	}
}
