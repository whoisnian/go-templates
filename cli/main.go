package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/whoisnian/go-templates/cli/global"
)

func main() {
	global.SetupConfig()
	global.SetupLogger()
	global.LOG.Debugf("use config: %+v", global.CFG)

	if global.CFG.Version {
		fmt.Printf("%s %s(%s)\n", global.AppName, global.Version, global.BuildTime)
		return
	}

	req, err := http.NewRequest(global.CFG.Method, global.CFG.URL, nil)
	if err != nil {
		global.LOG.Fatalf("NewRequest: %v", err)
	}

	if global.CFG.DryRun {
		global.LOG.Infof("[DRY-RUN] %s %s", req.Method, req.URL.String())
	} else {
		doRequest(req)
	}
}

func doRequest(req *http.Request) {
	start := time.Now()
	global.LOG.Debugf("request start")
	defer func() { global.LOG.Debugf("request end, cost %v", time.Since(start)) }()

	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		global.LOG.Fatalf("RoundTrip: %v", err)
	}
	defer res.Body.Close()
	global.LOG.Debugf("response %s", res.Status)

	io.Copy(os.Stdout, res.Body)
}
