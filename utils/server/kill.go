package server

import (
	"miningPoolCli/utils/mlog"
	"net/http"
	"os"
	"time"
)

func killHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" && r.Method != "POST" {
			http.Error(w, string(errJson.MethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		defer func() {
			go func() {
				time.Sleep(512 * time.Millisecond)
				mlog.LogOk("Killing by http request to /kill ...")
				os.Exit(1)
			}()
		}()

		w.Write([]byte(`{"status": true, "kill": "bye..."}`))
	}
}
