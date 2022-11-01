package template

import (
	"create-vbox-vm/internal/controller/configuration"
	"create-vbox-vm/internal/controller/vbox/common/command/ds"
	"create-vbox-vm/internal/controller/vbox/common/command/template/argument"
)

type CreateVmCommandFactory struct {
}

func NewCreateVmCommandFactory() CreateVmCommandFactory {
	return newCreateVmCommandFactory()
}

func newCreateVmCommandFactory() CreateVmCommandFactory {
	return CreateVmCommandFactory{}
}

func (factory CreateVmCommandFactory) CreateCommand(configuration configuration.Configuration) *ds.VBoxManageCommand {
	return &ds.VBoxManageCommand{
		Name: "createvm",
		Args: []string{
			argument.Arg("name", configuration.General.Name).AsString(),
			argument.Arg("basefolder", configuration.Create.Predefined.VmsHome).AsString(),
			argument.Arg("ostype", configuration.Create.Predefined.OsType).AsString(),
			argument.Opt("register").AsString(),
		},
	}
}
