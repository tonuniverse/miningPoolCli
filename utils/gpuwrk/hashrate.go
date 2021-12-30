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
	var genStats struct {
		Khs    int   `json:"khs"`    // khs | total hashrate
		Uptime int64 `json:"uptime"` // uptime
		Hs     []int `json:"hs"`     // hs | array of hashrates
	}

	for i, v := range *gpus {
		hsArr := config.MRgxKit.FindHashRate.FindAllString(v.ProcStderr.String(), -1)
		if len(hsArr) < 2 {
			return
		}

		hS := config.MRgxKit.FindDecimal.FindAllString(hsArr[len(hsArr)-1], -1)
		if len(hS) < 1 {
			return
		}

		sep := strings.Split(hS[0], ".")
		if len(sep) != 2 {
			return
		}

		perHashRate, err := strconv.Atoi(sep[0])
		if err != nil {
			return
		}

		(*gpus)[i].CurrentHashrate = perHashRate

		genStats.Khs += perHashRate
		genStats.Hs = append(genStats.Hs, perHashRate)

	}

	if config.UpdateStatsFile {
		genStats.Uptime = time.Now().Unix() - config.StartProgramTimestamp
		file, err := json.Marshal(genStats)
		if err != nil {
			mlog.LogFatalStackError(err)
		}
		_ = ioutil.WriteFile("stats.json", file, 0644)
	}

	mlog.LogInfo("Total hashrate: ~" + strconv.Itoa(genStats.Khs) + " Mh")
}
