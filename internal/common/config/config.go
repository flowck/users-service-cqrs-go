package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port           int    `envconfig:"PORT"`
	GrpcPort       int    `envconfig:"GRPC_PORT"`
	PsqlUri        string `envconfig:"GOOSE_DBSTRING"`
	ApplyPsqlSeeds string `envconfig:"APPLY_PSQL_SEEDS"`
}

func (c *Config) IsSeedsEnabled() bool {
	return c.ApplyPsqlSeeds == "enabled"
}

func New() *Config {
	cfg := &Config{}

	if os.Getenv("ENVIRONMENT") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("Unable to find a .env file")
		}
		log.Println("Config has been loaded from from .env because env is not set to production")
	}

	if err := envconfig.Process("", cfg); err != nil {
		panic(err)
	}

	return cfg
}
