package create

import (
	"create-vbox-vm/internal/controller/configuration"
	"create-vbox-vm/internal/controller/vbox/common/command"
	"create-vbox-vm/internal/controller/vbox/create/scenario"
)

type VBoxVmCreationFacade interface {
	Execute(configuration configuration.Configuration) error
}

type VBoxVmCreationFacadeImpl struct {
	vboxScenarioListExecutor     command.VBoxScenarioExecutor
	createScenarioProviderFacade scenario.CreateScenarioProviderFacade
}

func NewVBoxVmCreationFacade() VBoxVmCreationFacade {
	return newVBoxVmCreationFacade(command.NewVBoxScenarioExecutor(), scenario.NewCreateScenarioProviderFacade())
}

func newVBoxVmCreationFacade(vboxCommandListExecutor command.VBoxScenarioExecutor, createScenarioProviderFacade scenario.CreateScenarioProviderFacade) VBoxVmCreationFacade {
	return VBoxVmCreationFacadeImpl{vboxScenarioListExecutor: vboxCommandListExecutor, createScenarioProviderFacade: createScenarioProviderFacade}
}

func (facade VBoxVmCreationFacadeImpl) Execute(configuration configuration.Configuration) error {
	vboxVmCreateScenario := facade.createScenarioProviderFacade.GetVBoxVmCreateScenario(configuration)

	return facade.vboxScenarioListExecutor.Execute(vboxVmCreateScenario, configuration)
}
