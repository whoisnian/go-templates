package global

import (
	"cmp"
	"context"
	"io/fs"
	"slices"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/whoisnian/glb/logger"
	"github.com/whoisnian/go-templates/server/pkg/postgres"
	"github.com/whoisnian/go-templates/server/schema"
)

var DB *pgxpool.Pool

func SetupPostgres(ctx context.Context) {
	cfg, err := pgxpool.ParseConfig(CFG.DatabaseURI)
	if err != nil {
		LOG.Fatal(ctx, "pgxpool.ParseConfig", logger.Error(err))
	}
	if CFG.Debug {
		cfg.ConnConfig.Tracer = &postgres.Tracer{LOG: LOG}
	}

	DB, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		LOG.Fatal(ctx, "pgxpool.NewWithConfig", logger.Error(err))
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	if err = DB.Ping(timeoutCtx); err != nil {
		LOG.Fatal(ctx, "pgxpool.Ping", logger.Error(err))
	}
}

func InitDatabaseSchema(ctx context.Context) {
	files, err := schema.FS.ReadDir(".")
	if err != nil {
		LOG.Fatal(ctx, "schema.FS.ReadDir", logger.Error(err))
	}
	LOG.Debugf(ctx, "found %d schema files", len(files))

	slices.SortFunc(files, func(a, b fs.DirEntry) int { return cmp.Compare(a.Name(), b.Name()) })

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		LOG.Debugf(ctx, "read and exec '%s'...", file.Name())
		data, err := schema.FS.ReadFile(file.Name())
		if err != nil {
			LOG.Fatal(ctx, "schema.FS.ReadFile", logger.Error(err))
		}
		if _, err = DB.Exec(ctx, string(data)); err != nil {
			LOG.Fatal(ctx, "pgxpool.Exec", logger.Error(err))
		}
	}
}
