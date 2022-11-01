package predefined

import (
	"create-vbox-vm/internal/controller/configuration"
	"create-vbox-vm/internal/controller/vbox/common/command/ds"
	"create-vbox-vm/internal/controller/vbox/common/command/template"
	"create-vbox-vm/internal/controller/vbox/install/scenario/predefined/steps"
	"github.com/rs/zerolog/log"
)

var VBOX_COMMAND_FACTORY_LIST = []template.VBoxManageCommandFactory{
	template.NewUnattendedInstallCommandFactory(),
	steps.NewModifyVmBoot1CommandFactory(),
	steps.NewModifyVmBoot2CommandFactory(),
	steps.NewModifyVmBoot3CommandFactory(),
	steps.NewModifyVmBoot4CommandFactory(),
	template.NewStartVmCommandFactory(),
}

type PredefinedInstallScenarioProvider interface {
	GetVBoxManageCommands(configuration configuration.Configuration) []ds.VBoxManageCommand
}

type PredefinedInstallScenarioProviderImpl struct {
}

func NewPredefinedInstallScenarioProvider() PredefinedInstallScenarioProvider {
	return newPredefinedInstallScenarioProvider()
}

func newPredefinedInstallScenarioProvider() PredefinedInstallScenarioProvider {
	return PredefinedInstallScenarioProviderImpl{}
}

func (provider PredefinedInstallScenarioProviderImpl) GetVBoxManageCommands(configuration configuration.Configuration) []ds.VBoxManageCommand {
	var predefinedInstallCommands []ds.VBoxManageCommand

	if provider.isInstallationConfiguredCorrectly(configuration) {
		log.Info().Msgf("Running OS Installation phase")

		for _, commandFactory := range VBOX_COMMAND_FACTORY_LIST {
			vboxManageCommand := commandFactory.CreateCommand(configuration)

			if vboxManageCommand != nil {
				predefinedInstallCommands = append(predefinedInstallCommands, *vboxManageCommand)
			}
		}
	} else {
		log.Info().Msgf("Vm Configuration is incorrect - will skip predefined OS Installation phase")
	}

	return predefinedInstallCommands
}

func (provider PredefinedInstallScenarioProviderImpl) isInstallationConfiguredCorrectly(configuration configuration.Configuration) bool {
	if len(configuration.General.Name) <= 0 {
		log.Info().Msgf("Vm Name was not provided")
		return false
	} else if len(configuration.General.IsoFilePath) <= 0 {
		log.Info().Msgf("Iso File Path was not provided")
		return false
	} else if len(configuration.Install.Predefined.Credentials.User) <= 0 {
		log.Info().Msgf("User name was not provided")
		return false
	} else if len(configuration.Install.Predefined.Credentials.Password) <= 0 {
		log.Info().Msgf("User Password was not provided")
		return false
	} else {
		return true
	}
}
