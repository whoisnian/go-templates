package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/whoisnian/glb/logger"
	"github.com/whoisnian/go-templates/cli/global"
)

func main() {
	ctx := context.Background()
	global.SetupConfig(ctx)
	global.SetupLogger(ctx)
	global.LOG.Debugf(ctx, "use config: %+v", global.CFG)

	if global.CFG.Version {
		fmt.Printf("%s %s(%s)\n", global.AppName, global.Version, global.BuildTime)
		return
	}

	req, err := http.NewRequest(global.CFG.Method, global.CFG.URL, nil)
	if err != nil {
		global.LOG.Fatal(ctx, "http.NewRequest", logger.Error(err))
	}

	if global.CFG.DryRun {
		global.LOG.Infof(ctx, "[DRY-RUN] %s %s", req.Method, req.URL.String())
	} else {
		doRequest(ctx, req)
	}
}

func doRequest(ctx context.Context, req *http.Request) {
	start := time.Now()
	global.LOG.Debugf(ctx, "request start")
	defer func() { global.LOG.Debugf(ctx, "request end, cost %v", time.Since(start)) }()

	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		global.LOG.Fatal(ctx, "http.DefaultTransport.RoundTrip", logger.Error(err))
	}
	defer res.Body.Close()
	global.LOG.Debugf(ctx, "response %s", res.Status)

	io.Copy(os.Stdout, res.Body)
}
