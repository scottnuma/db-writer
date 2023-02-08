package main

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("db-writer is starting")

	externalDBURL, ok := os.LookupEnv("CONN_STRING")
	if !ok {
		log.Fatal().Msg("CONN_STRING must be set with the database's connection string")
	}

	db, err := sql.Open("postgres", externalDBURL)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open db")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to ping db")
	}

	_, err = db.Exec(setupTableStatement)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to setup stamps table")
		panic(err)
	}
	log.Info().Msg("successfully setup stamps table")

	for {
		_ = insertTime(db)
		time.Sleep(time.Second)
	}
}

const setupTableStatement = `
CREATE TABLE IF NOT EXISTS stamps (time timestamp);
`

const insertTimeStatement = `
INSERT into stamps (time) VALUES ($1);
`

func insertTime(db *sql.DB) error {
	_, err := db.Exec(insertTimeStatement, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("failed to insert time")
		return err
	}
	log.Info().Msg("inserted time")
	return nil
}
