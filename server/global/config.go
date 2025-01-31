package global

import (
	"context"

	"github.com/whoisnian/glb/config"
)

var CFG Config

type Config struct {
	Debug  bool   `flag:"d,false,Enable debug output"`
	LogFmt string `flag:"log,nano,Log output format, one of nano, text and json"`

	FirstRun bool `flag:"first,false,Initialize database and quit in the first run"`
	Version  bool `flag:"v,false,Show version and quit"`

	ListenAddr  string `flag:"l,127.0.0.1:9000,Server listen addr"`
	DatabaseURI string `flag:"db,postgresql://postgres@127.0.0.1/dbname,PostgreSQL database connection URI"`
}

func SetupConfig(_ context.Context) {
	_, err := config.FromCommandLine(&CFG)
	if err != nil {
		panic(err)
	}
}
