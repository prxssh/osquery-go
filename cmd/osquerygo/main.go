package main

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	initLogger()

	log.Info().Msg("osquerygo is up and running")
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
