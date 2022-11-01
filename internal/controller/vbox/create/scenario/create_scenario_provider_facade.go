package scenario

import (
	"create-vbox-vm/internal/controller/configuration"
	"create-vbox-vm/internal/controller/vbox/common/command/ds"
	"create-vbox-vm/internal/controller/vbox/common/scenario/raw"
	"create-vbox-vm/internal/controller/vbox/create/scenario/predefined"
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
