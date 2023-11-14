package homles

import (
	"mosn.io/holmes"
	mlog "mosn.io/pkg/log"
	"net/http"
	"testing"
	"time"
)

func TestCpu(t *testing.T) {
	http.HandleFunc("/cpuex", cpuex)
	go http.ListenAndServe(":10003", nil) //nolint:errcheck
	h, _ := holmes.New(
		holmes.WithCollectInterval("2s"),
		holmes.WithDumpPath("./tmp"),
		holmes.WithLogger(holmes.NewFileLog("./tmp/holmes.log", mlog.DEBUG)),
		holmes.WithCPUDump(20, 25, 80, time.Minute),
		holmes.WithTextDump(),
	)
	h.EnableCPUDump().Start()
	time.Sleep(time.Hour)
}
func cpuex(wr http.ResponseWriter, req *http.Request) {
	go func() {
		for {

		}
	}()
}
