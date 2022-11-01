package main

import (
	"create-vbox-vm/internal/controller/cli"
	"create-vbox-vm/internal/controller/configuration"
	"create-vbox-vm/internal/controller/logging"
	"create-vbox-vm/internal/controller/vbox"
	"github.com/rs/zerolog/log"
	"os"
)

var cliArgumentsReader = cli.NewCliArgumentsReader()
var configFilePathProvider = configuration.NewConfigurationFilePathProvider()
var configFileReader = configuration.NewConfigurationFileReader()
var VBoxVmSetup = vbox.NewVBoxVmSetup()

func main() {
	returnCode := run()

	os.Exit(returnCode)
}

func run() int {
	cliArguments, err := cliArgumentsReader.GetCliArguments()

	if err != nil {
		return 1
	} else if cliArguments == nil {
		println("Was not able to continue because of lack of arguments")
		return 1
	}

	logging.InitializeLogger(cliArguments.Debug)

	configFilePath := configFilePathProvider.GetConfigurationFilePath(cliArguments.ConfigPath)
	configuration := configFileReader.ReadConfiguration(configFilePath)

	err = VBoxVmSetup.PrepareNewVm(*configuration)

	if err != nil {
		log.Err(err).Msgf("Failed to create VM")
		return 1
	} else {
		return 0
	}
}
