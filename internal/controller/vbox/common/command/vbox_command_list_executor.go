package command

import (
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/common/command/ds"
)

type VBoxScenarioExecutor interface {
	Execute(vboxManageScenario []ds.VBoxManageCommand, configuration configuration.Configuration) error
}

type VBoxScenarioExecutorImpl struct {
	vboxCommandExecutor VBoxCommandExecutor
}

func NewVBoxScenarioExecutor() VBoxScenarioExecutor {
	return newVBoxScenarioExecutor(NewVBoxCommandExecutor())
}

func newVBoxScenarioExecutor(vboxCommandExecutor VBoxCommandExecutor) VBoxScenarioExecutor {
	return VBoxScenarioExecutorImpl{vboxCommandExecutor: vboxCommandExecutor}
}

func (executor VBoxScenarioExecutorImpl) Execute(vboxManageScenario []ds.VBoxManageCommand, configuration configuration.Configuration) error {

	for _, vboxManageCommand := range vboxManageScenario {
		err := executor.vboxCommandExecutor.Execute(configuration, vboxManageCommand)

		if err != nil {
			return err
		}
	}

	return nil
}
