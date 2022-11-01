package commons

import (
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
)

func GetApplicationRootPath() string {
	applicationPath, err := os.Executable()
	if err != nil {
		log.Panic().Err(err).Msg("Application root path not defined")
	}
	return filepath.Dir(applicationPath)
}
