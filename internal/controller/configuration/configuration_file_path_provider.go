package configuration

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"vbox-vm-setup/internal/controller/commons"
)

const DEFAULT_CONFIGURATION_FILE_NAME = "config.yml"

type ConfigurationFilePathProvider struct {
}

func NewConfigurationFilePathProvider() ConfigurationFilePathProvider {
	return newConfigurationFilePathProvider()
}

func newConfigurationFilePathProvider() ConfigurationFilePathProvider {
	return ConfigurationFilePathProvider{}
}

func (configurationFilePathProvider ConfigurationFilePathProvider) GetConfigurationFilePath(configurationFilePath string) string {
	if len(configurationFilePath) != 0 {
		return configurationFilePath
	} else {
		applicationRootPath := commons.GetApplicationRootPath()
		configFilePath := fmt.Sprintf("%s/%s", applicationRootPath, DEFAULT_CONFIGURATION_FILE_NAME)

		log.Info().Msgf("Assuming default config file location: %s", configFilePath)

		return configFilePath
	}
}
