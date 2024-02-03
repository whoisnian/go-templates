package main

import (
	"fmt"

	"github.com/whoisnian/go-templates/server/global"
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
}
