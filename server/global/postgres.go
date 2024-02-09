package global

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/whoisnian/go-templates/server/pkg/postgres"
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
