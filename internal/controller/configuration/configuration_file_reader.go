package configuration

import (
	"github.com/creasty/defaults"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
	"os"
)

type ConfigurationFileReader struct {
}

func NewConfigurationFileReader() ConfigurationFileReader {
	return newConfigurationFileReader()
}

func newConfigurationFileReader() ConfigurationFileReader {
	return ConfigurationFileReader{}
}

func (configurationFileReader ConfigurationFileReader) ReadConfiguration(configurationFilePath string) *Configuration {

	log.Info().Msgf("Reading configuration file")

	configurationFileContent, err := os.ReadFile(configurationFilePath)
	if err != nil {
		log.Panic().Err(err).Msgf("Error occurred while reading configuration file %s: %s", configurationFilePath, err)
	}

	configuration := Configuration{}

	if err := defaults.Set(&configuration); err != nil {
		log.Panic().Err(err).Msgf("Was not able to apply configuration defaults")
	}

	err = yaml.Unmarshal(configurationFileContent, &configuration)
	if err != nil {
		log.Panic().Err(err).Msgf("Error occurred while unmarshalling configuration file content %s: %s", configurationFilePath, err)
	}

	log.Info().Msgf("Configuration was read successfully")

	return &configuration
}
