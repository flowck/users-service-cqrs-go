package psql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/pressly/goose/v3"
)

func Connect(uri string) *sql.DB {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Panicf("could not open a connection to postgres: \nuri: %s\nerror: %v", uri, err)
	}

	if err = db.Ping(); err != nil {
		log.Panicf("unable to ping postgres: \nuri: %s\nerror: %v", uri, err)
	}

	return db
}

func ApplyPsqlMigrationsAndSeeds(db *sql.DB, seedsEnabled bool) {
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	workdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if err = goose.Up(db, fmt.Sprintf("%s/sql/migrations", workdir)); err != nil {
		panic(err)
	}

	if !seedsEnabled {
		return
	}

	if err = goose.Up(db, fmt.Sprintf("%s/sql/seeds", workdir)); err != nil {
		panic(err)
	}
}
