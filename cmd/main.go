package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"users-service-cqrs/internal/adapters"
	"users-service-cqrs/internal/app"
	"users-service-cqrs/internal/app/command"
	"users-service-cqrs/internal/app/query"
	"users-service-cqrs/internal/ports/http"

	"github.com/kelseyhightower/envconfig"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type config struct {
	Port    int    `envconfig:"PORT"`
	PsqlUri string `envconfig:"GOOSE_DBSTRING"`
}

func main() {
	cfg := getConfig()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)

	//
	// DB
	//
	db, err := sql.Open("postgres", cfg.PsqlUri)
	if err != nil {
		log.Fatalf("could not open a connection to db: %v", err)
	}

	applyPsqlMigrationsAndSeeds(db, true)

	//
	// Data Repositories
	//
	readUserRepo := adapters.NewPsqlReadUserRepo(db)
	inMemWriteUserRepo := adapters.NewInMemoryWriteUserRepository()

	//
	// Assemble the application
	//
	application := &app.App{
		Commands: &app.Commands{
			BlockUser:   command.NewBlockUserHandler(inMemWriteUserRepo),
			UnBlockUser: command.NewUnblockUserHandler(inMemWriteUserRepo),
		},
		Queries: &app.Queries{
			AllUsers: query.NewAllBlockedUsersHandler(readUserRepo),
			OneUser:  query.NewOneUserHandler(readUserRepo),
		},
	}

	httpServer := http.NewServer(ctx, cfg.Port, application)
	httpServer.Start()

	<-done
	log.Println("Stopping the service gracefully")

	ctx, cancel = context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	httpServer.Stop(ctx)
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

func getConfig() *config {
	cfg := &config{}

	if err := envconfig.Process("", cfg); err != nil {
		panic(err)
	}

	return cfg
}
