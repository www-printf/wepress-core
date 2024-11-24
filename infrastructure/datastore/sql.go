package datastore

import (
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbDSN string

func ProvideDatabase(dbDSN DbDSN) *gorm.DB {
	db, err := gorm.Open(postgres.Open(string(dbDSN)), &gorm.Config{})
	if err != nil {
		log.Panic().Msgf("Error when create DB: %v", err)
	}

	normalDB, err := db.DB()

	if err != nil {
		log.Panic().Msgf("Error when checking DB: %v", err)
	}

	if err := normalDB.Ping(); err != nil {
		log.Panic().Msgf("Error when ping DB: %v", err)
	}

	normalDB.SetConnMaxIdleTime(time.Minute * 5)
	normalDB.SetMaxIdleConns(2)
	normalDB.SetMaxOpenConns(5)

	return db
}
