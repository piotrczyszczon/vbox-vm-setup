package predefined

import (
	"create-vbox-vm/internal/controller/configuration"
	"create-vbox-vm/internal/controller/vbox/common/command/ds"
	"create-vbox-vm/internal/controller/vbox/common/command/template"
	"create-vbox-vm/internal/controller/vbox/create/scenario/predefined/steps"
	"github.com/rs/zerolog/log"
)

var VBOX_COMMAND_FACTORY_LIST = []template.VBoxManageCommandFactory{
	template.NewCreateVmCommandFactory(),
	steps.NewModifyVmMemoryCommandFactory(),
	steps.NewModifyVmCpusCommandFactory(),
	steps.NewModifyVmVirtualMemoryCommandFactory(),
	steps.NewModifyVmGraphicsControllerCommandFactory(),
	steps.NewModifyVmRtcUseUtcCommandFactory(),
	steps.NewModifyVmClipboardModeCommandFactory(),
	steps.NewModifyVmDragAndDropCommandFactory(),
	steps.NewModifyVmPaeCommandFactory(),
	steps.NewModifyVmNatRuleCommandFactory(),
	steps.NewStorageCtlSataCommandFactory(),
	template.NewCreatMediumCommandFactory(),
	steps.NewStorageAttachHddCommandFactory(),
	steps.NewStorageCtlIdeCommandFactory(),
	steps.NewStorageAttachDvdDriveCommandFactory(),
}

type PredefinedCreateScenarioProvider interface {
	GetVBoxManageCommands(configuration configuration.Configuration) []ds.VBoxManageCommand
}

type PredefinedCreateScenarioProviderImpl struct {
}

func NewPredefinedCreateScenarioProvider() PredefinedCreateScenarioProvider {
	return newPredefinedCreateScenarioProvider()
}

func newPredefinedCreateScenarioProvider() PredefinedCreateScenarioProvider {
	return PredefinedCreateScenarioProviderImpl{}
}

func (provider PredefinedCreateScenarioProviderImpl) GetVBoxManageCommands(configuration configuration.Configuration) []ds.VBoxManageCommand {
	var PredefinedCreateCommands []ds.VBoxManageCommand

	if provider.isVmConfiguredCorrectly(configuration) {
		log.Info().Msgf("Running VM Creation phase")

		for _, commandFactory := range VBOX_COMMAND_FACTORY_LIST {
			vboxManageCommand := commandFactory.CreateCommand(configuration)

			if vboxManageCommand != nil {
				PredefinedCreateCommands = append(PredefinedCreateCommands, *vboxManageCommand)
			}
		}
	} else {
		log.Info().Msgf("Vm Configuration is incorrect - will skip predefined VM Creation phase")
	}

	return PredefinedCreateCommands
}

func (provider PredefinedCreateScenarioProviderImpl) isVmConfiguredCorrectly(configuration configuration.Configuration) bool {
	if len(configuration.General.Name) <= 0 {
		log.Info().Msgf("Vm Name was not provided")
		return false
	} else if len(configuration.Create.Predefined.VmsHome) <= 0 {
		log.Info().Msgf("Vms Home was not provided")
		return false
	} else if len(configuration.General.IsoFilePath) <= 0 {
		log.Info().Msgf("Iso File Path was not provided")
		return false
	} else {
		return true
	}
}
