package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"users-service-cqrs/internal/app"
	"users-service-cqrs/internal/common/config"
	"users-service-cqrs/internal/common/psql"
	"users-service-cqrs/internal/ports/grpc_port"
	"users-service-cqrs/internal/ports/http"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.New()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)

	db := psql.Connect(cfg.PsqlUri)
	psql.ApplyPsqlMigrationsAndSeeds(db, cfg.IsSeedsEnabled())

	application := app.New(db)

	//
	// Http server
	//
	httpServer := http.NewServer(ctx, cfg.Port, application)
	httpServer.Start()
	log.Printf("The http server is running at http://localhost:%d\n", cfg.Port)

	//
	// Grpc server
	//
	grpcServer := grpc_port.NewServer()
	grpcServer.Start(cfg.GrpcPort, application)

	<-done
	log.Println("Stopping the service gracefully")

	ctx, cancel = context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	httpServer.Stop(ctx)
	grpcServer.Stop()
}
