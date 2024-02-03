package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/whoisnian/glb/util/osutil"
	"github.com/whoisnian/go-templates/server/global"
	"github.com/whoisnian/go-templates/server/router"
)

func main() {
	global.SetupConfig()
	global.SetupLogger()
	global.LOG.Debugf("use config: %+v", global.CFG)

	if global.CFG.Version {
		fmt.Printf("%s %s(%s)\n", global.AppName, global.Version, global.BuildTime)
		return
	}
	if global.CFG.FirstRun {
		// TODO: initialize database
		return
	}

	server := &http.Server{Addr: global.CFG.ListenAddr, Handler: router.Setup()}
	go func() {
		global.LOG.Infof("Service started: http://%s", global.CFG.ListenAddr)
		if err := server.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
			global.LOG.Warn("Service shutting down")
		} else if err != nil {
			global.LOG.Fatal(err.Error())
		}
	}()

	osutil.WaitForStop()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		global.LOG.Warn(err.Error())
	}
}
