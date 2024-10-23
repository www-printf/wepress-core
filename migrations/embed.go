package migrations

import (
	"database/sql"
	"embed"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//go:embed sql/*
var fs embed.FS

type MLog struct {
	log zerolog.Logger
}

func (MLog) Verbose() bool { return false }
func (l MLog) Printf(format string, v ...interface{}) {
	l.log.Info().Msgf(format, v...)
}

func RunAutoMigrate(db *sql.DB) {
	d, err := iofs.New(fs, "sql")
	if err != nil {
		log.Fatal().Msgf("auto migration - init iofs: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal().Msgf("auto migration - init postgres driver: %v", err)
	}

	m, err := migrate.NewWithInstance("iofs", d, "postgres", driver)
	if err != nil {
		log.Fatal().Msgf("auto migration - init migrate: %v", err)
	}

	defer m.Close()
	m.Log = &MLog{log: log.Logger}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Msgf("auto migration - up: %v", err)
	}
	dbversion, dirty, err := m.Version()
	if err != nil {
		log.Fatal().Msgf("auto migration - error get db version: %v", err)
	}

	log.Info().Msgf("auto migration - db version: %v, dirty: %v", dbversion, dirty)
}
