package server

import (
	"io/ioutil"
	"miningPoolCli/config"
	"miningPoolCli/utils/gpuwrk"
	"miningPoolCli/utils/mlog"
	"net"
	"net/http"
	"strconv"
)

func Entrypoint(gpuData *[]gpuwrk.GpuGoroutine) {
	http.HandleFunc("/stat", statHandler(gpuData))

	if config.NetSrv.HandleKill {
		http.HandleFunc("/kill", killHandler(gpuData))
		mlog.LogInfo("Set kill http handler at /kill")
	}

	listener, err := net.Listen("tcp", config.NetSrv.Host+":0")
	if err != nil {
		mlog.LogError("Failed to get tcp")
		mlog.LogFatalStackError(err)
	}
	hostPort := config.NetSrv.Host + ":" + strconv.Itoa(listener.Addr().(*net.TCPAddr).Port)

	if err := ioutil.WriteFile(config.NetSrv.HostFileName, []byte(hostPort), 0644); err != nil {
		mlog.LogFatalStackError(err)
	}
	mlog.LogInfo("Server addr saved to: " + config.NetSrv.HostFileName)

	mlog.LogInfo("Server at: " + hostPort)
	if err := http.Serve(listener, nil); err != nil {
		mlog.LogFatalStackError(err)
	}
}
