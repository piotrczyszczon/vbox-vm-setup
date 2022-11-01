package scenario

import (
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/common/command/ds"
	"vbox-vm-setup/internal/controller/vbox/common/scenario/raw"
	"vbox-vm-setup/internal/controller/vbox/create/scenario/predefined"
)

type CreateScenarioProviderFacade interface {
	GetVBoxVmCreateScenario(configuration configuration.Configuration) []ds.VBoxManageCommand
}

type CreateScenarioProviderFacadeImpl struct {
	predefinedCreateScenarioProvider predefined.PredefinedCreateScenarioProvider
	rawCreateScenarioProvider        raw.RawScenarioProvider
}

func NewCreateScenarioProviderFacade() CreateScenarioProviderFacade {
	return newCreateScenarioProviderFacade(predefined.NewPredefinedCreateScenarioProvider(), raw.NewRawScenarioProvider())
}

func newCreateScenarioProviderFacade(predefinedCreateScenarioProvider predefined.PredefinedCreateScenarioProvider, rawCreateScenarioProvider raw.RawScenarioProvider) CreateScenarioProviderFacade {
	return CreateScenarioProviderFacadeImpl{predefinedCreateScenarioProvider: predefinedCreateScenarioProvider, rawCreateScenarioProvider: rawCreateScenarioProvider}
}

func (facade CreateScenarioProviderFacadeImpl) GetVBoxVmCreateScenario(configuration configuration.Configuration) []ds.VBoxManageCommand {
	var vboxManageCommands []ds.VBoxManageCommand

	vboxManageCommands = append(vboxManageCommands, facade.predefinedCreateScenarioProvider.GetVBoxManageCommands(configuration)...)
	vboxManageCommands = append(vboxManageCommands, facade.rawCreateScenarioProvider.GetVBoxManageCommands(configuration.Create.Raw)...)

	return vboxManageCommands
}
