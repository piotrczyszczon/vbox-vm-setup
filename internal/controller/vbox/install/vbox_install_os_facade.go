package install

import (
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/common/command"
	"vbox-vm-setup/internal/controller/vbox/install/scenario"
)

type VBoxInstallOSFacade interface {
	Execute(configuration configuration.Configuration) error
}

type VBoxInstallOSFacadeImpl struct {
	vboxCommandListExecutor       command.VBoxScenarioExecutor
	installScenarioProviderFacade scenario.InstallScenarioProviderFacade
}

func NewVBoxInstallOSFacade() VBoxInstallOSFacade {
	return newVBoxInstallOSFacade(command.NewVBoxScenarioExecutor(), scenario.NewInstallScenarioProviderFacade())
}

func newVBoxInstallOSFacade(vboxCommandListExecutor command.VBoxScenarioExecutor, installScenarioProviderFacade scenario.InstallScenarioProviderFacade) VBoxInstallOSFacade {
	return VBoxInstallOSFacadeImpl{vboxCommandListExecutor: vboxCommandListExecutor, installScenarioProviderFacade: installScenarioProviderFacade}
}

func (facade VBoxInstallOSFacadeImpl) Execute(configuration configuration.Configuration) error {
	vboxVmCreateScenario := facade.installScenarioProviderFacade.GetVBoxVmInstallScenario(configuration)

	return facade.vboxCommandListExecutor.Execute(vboxVmCreateScenario, configuration)
}
