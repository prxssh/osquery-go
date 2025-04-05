package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type env struct {
	Postgres struct {
		User     string `required:"true" split_words:"true"`
		Password string `required:"true" split_words:"true"`
		Host     string `required:"true" split_words:"true"`
		Port     string `required:"true" split_words:"true"`
		Dbname   string `required:"true" split_words:"true"`
	}
	GooseMigrationDir     string `required:"true" split_words:"true"`
	OsquerySocketFilePath string `required:"true" split_words:"true"`
}

var Env env

func Init() {
	if err := envconfig.Process("", &Env); err != nil {
		log.Fatal().Err(err).Msg("config: failed")
	}

	log.Info().Msg("config: loaded")
}
