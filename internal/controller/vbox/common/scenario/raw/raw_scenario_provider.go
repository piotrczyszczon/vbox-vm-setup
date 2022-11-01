package raw

import (
	"create-vbox-vm/internal/controller/configuration"
	"create-vbox-vm/internal/controller/vbox/common/command/ds"
	"github.com/rs/zerolog/log"
)

type RawScenarioProvider interface {
	GetVBoxManageCommands(rawScenario configuration.Raw) []ds.VBoxManageCommand
}

type RawScenarioProviderImpl struct {
}

func NewRawScenarioProvider() RawScenarioProvider {
	return newRawScenarioProvider()
}

func newRawScenarioProvider() RawScenarioProvider {
	return RawScenarioProviderImpl{}
}

func (provider RawScenarioProviderImpl) GetVBoxManageCommands(rawScenario configuration.Raw) []ds.VBoxManageCommand {
	var rawCommands []ds.VBoxManageCommand

	if len(rawScenario.Commands) > 0 {
		log.Info().Msgf("Executing [%d] raw commands", len(rawScenario.Commands))

		for _, rawCommand := range rawScenario.Commands {
			rawCommands = append(rawCommands, ds.VBoxManageCommand{Name: rawCommand.Name, Args: rawCommand.Args})
		}
	} else {
		log.Info().Msgf("Not found any Raw commands - will skip Raw phase")
	}

	return rawCommands
}
