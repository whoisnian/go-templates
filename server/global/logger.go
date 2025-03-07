package global

import (
	"context"
	"os"

	"github.com/whoisnian/glb/ansi"
	"github.com/whoisnian/glb/logger"
	"github.com/whoisnian/go-templates/server/pkg/tracer"
)

var LOG *logger.Logger

func SetupLogger(_ context.Context) {
	opts := logger.Options{
		Level:     logger.LevelInfo,
		Colorful:  false,
		AddSource: true,
	}
	if CFG.Debug {
		opts.Level = logger.LevelDebug
		opts.Colorful = ansi.IsSupported(os.Stderr.Fd())
	}

	switch CFG.LogFmt {
	case "nano":
		LOG = logger.New(&tracer.WrapHandler{
			Handler: logger.NewNanoHandler(os.Stderr, opts),
		})
	case "text":
		LOG = logger.New(&tracer.WrapHandler{
			Handler: logger.NewTextHandler(os.Stderr, opts),
		})
	case "json":
		LOG = logger.New(&tracer.WrapHandler{
			Handler: logger.NewJsonHandler(os.Stderr, opts),
		})
	default:
		panic("unknown log format")
	}
}
