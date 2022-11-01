package commons

import (
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
)

func GetApplicationName() string {
	applicationPath, err := os.Executable()
	if err != nil {
		log.Panic().Err(err).Msg("Application name not defined")
	}
	return filepath.Base(applicationPath)
}
