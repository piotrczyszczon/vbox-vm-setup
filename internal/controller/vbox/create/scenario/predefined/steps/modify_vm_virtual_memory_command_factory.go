package steps

import (
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/common/command/ds"
	"vbox-vm-setup/internal/controller/vbox/common/command/template"
)

type ModifyVmVirtualMemoryCommandFactory struct {
	modifyVmCommandFactory template.ModifyVmCommandFactory
}

func NewModifyVmVirtualMemoryCommandFactory() ModifyVmVirtualMemoryCommandFactory {
	return newModifyVmVirtualMemoryCommandFactory(template.NewModifyVmCommandFactory())
}

func newModifyVmVirtualMemoryCommandFactory(modifyVmCommandFactory template.ModifyVmCommandFactory) ModifyVmVirtualMemoryCommandFactory {
	return ModifyVmVirtualMemoryCommandFactory{modifyVmCommandFactory: modifyVmCommandFactory}
}

func (factory ModifyVmVirtualMemoryCommandFactory) CreateCommand(configuration configuration.Configuration) *ds.VBoxManageCommand {
	return factory.modifyVmCommandFactory.CreateCommand(
		"vram", configuration.Create.Predefined.VirtualMemory, configuration,
	)
}
