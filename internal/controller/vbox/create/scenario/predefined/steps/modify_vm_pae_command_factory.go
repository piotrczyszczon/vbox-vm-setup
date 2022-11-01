package steps

import (
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/common/command/ds"
	"vbox-vm-setup/internal/controller/vbox/common/command/template"
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
