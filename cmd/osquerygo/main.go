package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/osquery/osquery-go"
	"github.com/pressly/goose/v3"

	"github.com/prxssh/osquery-go/api"
	"github.com/prxssh/osquery-go/config"
	"github.com/prxssh/osquery-go/config/postgres"
	"github.com/prxssh/osquery-go/internal/repo"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	initLogger()
	config.Init()
	runGooseMigrations()

	dbClient := initPostgres()
	repo := repo.NewRepo(dbClient)

	if err := api.StartServer(repo); err != nil {
		log.Fatal().Msgf("Failed to start server: %v", err)
	}
}

func initLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}

	log.Logger = log.With().Caller().Logger()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func initPostgres() *postgres.PostgresClient {
	client, err := postgres.Init()
	if err != nil {
		log.Fatal().Msgf("postgres: %v", err)
	}
	log.Info().Msg("postgres: connected successfully!")
	return client
}

func runGooseMigrations() {
	dbstring := fmt.Sprintf(
		"host=%s "+"port=%s "+"user=%s "+"password=%s "+"dbname=%s "+"sslmode=disable",
		config.Env.Postgres.Host,
		config.Env.Postgres.Port,
		config.Env.Postgres.User,
		config.Env.Postgres.Password,
		config.Env.Postgres.Dbname,
	)

	db, err := goose.OpenDBWithDriver("postgres", dbstring)
	if err != nil {
		log.Fatal().Msgf("goose: %v", err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal().Msgf("goose: %v", err)
	}

	if err := goose.Up(db, config.Env.GooseMigrationDir); err != nil {
		log.Fatal().Msgf("goose: %v", err)
	}

	log.Info().Msg("goose: completed successfully!")
}

func newOsqueryClient() osquery.ExtensionManager {
	fmt.Println("osquery client", config.Env.OsquerySocketFilePath)
	client, err := osquery.NewClient(
		config.Env.OsquerySocketFilePath,
		10*time.Second,
	)
	if err != nil {
		log.Fatal().Msgf("osquery: error creating osquery client: %v", err)
	}
	return client
}
