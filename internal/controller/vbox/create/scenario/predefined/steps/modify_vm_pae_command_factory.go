package steps

import (
	"create-vbox-vm/internal/controller/configuration"
	"create-vbox-vm/internal/controller/vbox/common/command/ds"
	"create-vbox-vm/internal/controller/vbox/common/command/template"
)

type ModifyVmPaeCommandFactory struct {
	modifyVmCommandFactory template.ModifyVmCommandFactory
}

func NewModifyVmPaeCommandFactory() ModifyVmPaeCommandFactory {
	return newModifyVmPaeCommandFactory(template.NewModifyVmCommandFactory())
}

func newModifyVmPaeCommandFactory(modifyVmCommandFactory template.ModifyVmCommandFactory) ModifyVmPaeCommandFactory {
	return ModifyVmPaeCommandFactory{modifyVmCommandFactory: modifyVmCommandFactory}
}

func (factory ModifyVmPaeCommandFactory) CreateCommand(configuration configuration.Configuration) *ds.VBoxManageCommand {
	return factory.modifyVmCommandFactory.CreateCommand(
		"pae", configuration.Create.Predefined.Pae, configuration,
	)
}
