package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"users-service-cqrs/internal/adapters"
	"users-service-cqrs/internal/app"
	"users-service-cqrs/internal/app/command"
	"users-service-cqrs/internal/app/query"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//
	// DB
	//
	db, err := sql.Open("postgres", os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		log.Fatalf("could not open a connection to db: %v", err)
	}

	applyPsqlMigrationsAndSeeds(db, true)

	//
	// Data Repositories
	//
	inMemReadUserRepo := adapters.NewInMemoryReadUserRepository()
	inMemWriteUserRepo := adapters.NewInMemoryWriteUserRepository()

	//
	// Assemble the application
	//
	application := app.App{
		Commands: &app.Commands{
			BlockUser:   command.NewBlockUserHandler(inMemWriteUserRepo),
			UnBlockUser: command.NewUnblockUserHandler(inMemWriteUserRepo),
		},
		Queries: &app.Queries{
			AllBlockedUser: query.NewAllBlockedUsersHandler(inMemReadUserRepo),
		},
	}

	fmt.Println(application.Queries.AllBlockedUser.Handle(ctx, query.AllBlockedUsers{}))

	server := http.Server{
		Addr:              fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler:           nil,
		ReadHeaderTimeout: time.Second * 2,
		IdleTimeout:       time.Second * 2,
		ReadTimeout:       time.Second * 2,
		WriteTimeout:      time.Second * 2,
	}

	if err = server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func applyPsqlMigrationsAndSeeds(db *sql.DB, seedsEnabled bool) {
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
