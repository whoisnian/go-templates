package global

import (
	"cmp"
	"context"
	"io/fs"
	"slices"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/whoisnian/go-templates/server/pkg/postgres"
	"github.com/whoisnian/go-templates/server/schema"
)

var DB *pgxpool.Pool

func SetupPostgres() {
	cfg, err := pgxpool.ParseConfig(CFG.DatabaseURI)
	if err != nil {
		LOG.Fatal(err.Error())
	}
	if CFG.Debug {
		cfg.ConnConfig.Tracer = &postgres.Tracer{LOG: LOG}
	}

	DB, err = pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		LOG.Fatal(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err = DB.Ping(ctx); err != nil {
		LOG.Fatal(err.Error())
	}
}

func InitDatabaseSchema() {
	files, err := schema.FS.ReadDir(".")
	if err != nil {
		LOG.Fatal(err.Error())
	}
	LOG.Debugf("found %d schema files", len(files))

	slices.SortFunc(files, func(a, b fs.DirEntry) int { return cmp.Compare(a.Name(), b.Name()) })

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		LOG.Debugf("read and exec '%s'...", file.Name())
		data, err := schema.FS.ReadFile(file.Name())
		if err != nil {
			LOG.Fatal(err.Error())
		}
		if _, err = DB.Exec(context.Background(), string(data)); err != nil {
			LOG.Fatal(err.Error())
		}
	}
}
