package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port                int    `envconfig:"PORT"`
	GrpcPort            int    `envconfig:"GRPC_PORT"`
	PsqlUri             string `envconfig:"GOOSE_DBSTRING"`
	ApplyPsqlMigrations bool   `envconfig:"APPLY_PSQL_MIGRATIONS"`
}

func New() *Config {
	cfg := &Config{}

	if os.Getenv("ENVIRONMENT") != "production" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
		log.Println("Config has been loaded from from .env because env is not set to production")
	}

	if err := envconfig.Process("", cfg); err != nil {
		panic(err)
	}

	return cfg
}
