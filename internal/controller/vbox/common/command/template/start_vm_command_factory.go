package template

import (
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/common/command/ds"
	"vbox-vm-setup/internal/controller/vbox/common/command/template/argument"
)

type StartVmCommandFactory struct {
}

func NewStartVmCommandFactory() StartVmCommandFactory {
	return newStartVmCommandFactory()
}

func newStartVmCommandFactory() StartVmCommandFactory {
	return StartVmCommandFactory{}
}

func (provider StartVmCommandFactory) CreateCommand(configuration configuration.Configuration) *ds.VBoxManageCommand {
	return &ds.VBoxManageCommand{
		Name: "startvm",
		Args: []string{
			argument.Param(configuration.General.Name).AsString(),
		},
	}
}
