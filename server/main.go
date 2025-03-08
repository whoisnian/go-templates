package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/whoisnian/glb/logger"
	"github.com/whoisnian/glb/util/osutil"
	"github.com/whoisnian/go-templates/server/global"
	"github.com/whoisnian/go-templates/server/router"
)

func main() {
	ctx := context.Background()
	global.SetupConfig(ctx)
	global.SetupLogger(ctx)
	global.LOG.Debugf(ctx, "use config: %+v", global.CFG)

	if global.CFG.Version {
		fmt.Printf("%s version %s built with %s at %s\n", global.AppName, global.Version, runtime.Version(), global.BuildTime)
		return
	}

	global.SetupPostgres(ctx)
	defer global.DB.Close()
	global.LOG.Debug(ctx, "connect to postgresql successfully")
	if global.CFG.FirstRun {
		global.LOG.Info(ctx, "initialize database schema in the first run")
		global.InitDatabaseSchema(ctx)
		return
	}

	server := &http.Server{Addr: global.CFG.ListenAddr, Handler: router.Setup(ctx)}
	go func() {
		global.LOG.Infof(ctx, "service started: http://%s", global.CFG.ListenAddr)
		if err := server.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
			global.LOG.Warn(ctx, "service shutting down")
		} else if err != nil {
			global.LOG.Fatal(ctx, "service start", logger.Error(err))
		}
	}()

	osutil.WaitForStop()

	shutdownCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		global.LOG.Warn(ctx, "service stop", logger.Error(err))
	}
}
