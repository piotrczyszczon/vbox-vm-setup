package logging

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func InitializeLogger(debug bool) {
	initializeLogger(debug)
}

func initializeLogger(debug bool) {
	zerolog.SetGlobalLevel(getGlobalLevel(debug))

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: isColorDisabled()})
}

func getGlobalLevel(debug bool) zerolog.Level {
	if debug {
		return zerolog.DebugLevel
	} else {
		return zerolog.InfoLevel
	}
}
