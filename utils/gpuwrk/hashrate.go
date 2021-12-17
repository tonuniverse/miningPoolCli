package gpuwrk

import (
	"encoding/json"
	"io/ioutil"
	"miningPoolCli/config"
	"miningPoolCli/utils/mlog"
	"strconv"
	"strings"
	"time"
)

func CalcHashrate(gpus *[]GpuGoroutine) {
	var totalHashRate int

	for i, v := range *gpus {
		hsArr := config.MRgxKit.FindHashRate.FindAllString(v.ProcStderr.String(), -1)
		if len(hsArr) < 2 {
			return
		}

		hS := config.MRgxKit.FindDecimal.FindAllString(hsArr[len(hsArr)-1], -1)

		sep := strings.Split(hS[0], ".")
		if len(sep) != 2 {
			return
		}

		perHashRate, err := strconv.Atoi(sep[0])
		if err != nil {
			return
		}

		(*gpus)[i].CurrentHashrate = perHashRate

		totalHashRate += perHashRate
	}

	if config.UpdateStatsFile {
		genStats := struct {
			HashRate int   `json:"hash_rate"`
			Uptime   int64 `json:"uptime"`
		}{
			HashRate: totalHashRate,
			Uptime:   time.Now().Unix() - config.StartProgramTimestamp,
		}

		file, err := json.Marshal(genStats)
		if err != nil {
			mlog.LogFatalStackError(err)
		}
		_ = ioutil.WriteFile("stats.json", file, 0644)
	}

	mlog.LogInfo("Total hashrate: ~" + strconv.Itoa(totalHashRate) + " Mh")
}
