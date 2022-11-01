package scenario

import (
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/common/command/ds"
	"vbox-vm-setup/internal/controller/vbox/common/scenario/raw"
	"vbox-vm-setup/internal/controller/vbox/install/scenario/predefined"
)

type InstallScenarioProviderFacade interface {
	GetVBoxVmInstallScenario(configuration configuration.Configuration) []ds.VBoxManageCommand
}

type InstallScenarioProviderFacadeImpl struct {
	predefinedInstallScenarioProvider predefined.PredefinedInstallScenarioProvider
	rawScenarioProvider               raw.RawScenarioProvider
}

func NewInstallScenarioProviderFacade() InstallScenarioProviderFacade {
	return newInstallScenarioProviderFacade(predefined.NewPredefinedInstallScenarioProvider(), raw.NewRawScenarioProvider())
}

func newInstallScenarioProviderFacade(predefinedInstallScenarioProvider predefined.PredefinedInstallScenarioProvider, rawScenarioProvider raw.RawScenarioProvider) InstallScenarioProviderFacade {
	return InstallScenarioProviderFacadeImpl{predefinedInstallScenarioProvider: predefinedInstallScenarioProvider, rawScenarioProvider: rawScenarioProvider}
}

func (facade InstallScenarioProviderFacadeImpl) GetVBoxVmInstallScenario(configuration configuration.Configuration) []ds.VBoxManageCommand {
	var vboxManageCommands []ds.VBoxManageCommand

	vboxManageCommands = append(vboxManageCommands, facade.predefinedInstallScenarioProvider.GetVBoxManageCommands(configuration)...)
	vboxManageCommands = append(vboxManageCommands, facade.rawScenarioProvider.GetVBoxManageCommands(configuration.Install.Raw)...)

	return vboxManageCommands
}
