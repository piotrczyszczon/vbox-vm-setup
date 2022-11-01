package steps

import (
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/common/command/ds"
	"vbox-vm-setup/internal/controller/vbox/common/command/template"
)

type ModifyVmMemoryCommandFactory struct {
	modifyVmCommandFactory template.ModifyVmCommandFactory
}

func NewModifyVmMemoryCommandFactory() ModifyVmMemoryCommandFactory {
	return newModifyVmMemoryCommandFactory(template.NewModifyVmCommandFactory())
}

func newModifyVmMemoryCommandFactory(modifyVmCommandFactory template.ModifyVmCommandFactory) ModifyVmMemoryCommandFactory {
	return ModifyVmMemoryCommandFactory{modifyVmCommandFactory: modifyVmCommandFactory}
}

func (factory ModifyVmMemoryCommandFactory) CreateCommand(configuration configuration.Configuration) *ds.VBoxManageCommand {
	return factory.modifyVmCommandFactory.CreateCommand(
		"memory", configuration.Create.Predefined.Memory, configuration,
	)
}
