package template

import (
	"create-vbox-vm/internal/controller/configuration"
	"create-vbox-vm/internal/controller/vbox/common/command/ds"
	"create-vbox-vm/internal/controller/vbox/common/command/template/argument"
)

type ModifyVmCommandFactory struct {
}

func NewModifyVmCommandFactory() ModifyVmCommandFactory {
	return newModifyVmCommandFactory()
}

func newModifyVmCommandFactory() ModifyVmCommandFactory {
	return ModifyVmCommandFactory{}
}

func (provider ModifyVmCommandFactory) CreateCommand(paramName string, paramValue string, configuration configuration.Configuration) *ds.VBoxManageCommand {
	return &ds.VBoxManageCommand{
		Name: "modifyvm",
		Args: []string{
			argument.Param(configuration.General.Name).AsString(),
			argument.Arg(paramName, paramValue).AsString(),
		},
	}
}
