package template

import (
	"create-vbox-vm/internal/controller/configuration"
	"create-vbox-vm/internal/controller/vbox/common/command/ds"
	"create-vbox-vm/internal/controller/vbox/common/command/template/argument"
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
