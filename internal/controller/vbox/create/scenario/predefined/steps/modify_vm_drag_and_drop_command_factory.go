package steps

import (
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/common/command/ds"
	"vbox-vm-setup/internal/controller/vbox/common/command/template"
)

type ModifyVmDragAndDropCommandFactory struct {
	modifyVmCommandFactory template.ModifyVmCommandFactory
}

func NewModifyVmDragAndDropCommandFactory() ModifyVmDragAndDropCommandFactory {
	return newModifyVmDragAndDropCommandFactory(template.NewModifyVmCommandFactory())
}

func newModifyVmDragAndDropCommandFactory(modifyVmCommandFactory template.ModifyVmCommandFactory) ModifyVmDragAndDropCommandFactory {
	return ModifyVmDragAndDropCommandFactory{modifyVmCommandFactory: modifyVmCommandFactory}
}

func (factory ModifyVmDragAndDropCommandFactory) CreateCommand(configuration configuration.Configuration) *ds.VBoxManageCommand {
	return factory.modifyVmCommandFactory.CreateCommand(
		"draganddrop", configuration.Create.Predefined.DragAndDrop, configuration,
	)
}
