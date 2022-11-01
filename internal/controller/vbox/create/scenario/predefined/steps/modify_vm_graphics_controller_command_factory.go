package steps

import (
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/common/command/ds"
	"vbox-vm-setup/internal/controller/vbox/common/command/template"
)

type ModifyVmGraphicsControllerCommandFactory struct {
	modifyVmCommandFactory template.ModifyVmCommandFactory
}

func NewModifyVmGraphicsControllerCommandFactory() ModifyVmGraphicsControllerCommandFactory {
	return newModifyVmGraphicsControllerCommandFactory(template.NewModifyVmCommandFactory())
}

func newModifyVmGraphicsControllerCommandFactory(modifyVmCommandFactory template.ModifyVmCommandFactory) ModifyVmGraphicsControllerCommandFactory {
	return ModifyVmGraphicsControllerCommandFactory{modifyVmCommandFactory: modifyVmCommandFactory}
}

func (factory ModifyVmGraphicsControllerCommandFactory) CreateCommand(configuration configuration.Configuration) *ds.VBoxManageCommand {
	return factory.modifyVmCommandFactory.CreateCommand(
		"graphicscontroller", configuration.Create.Predefined.GraphicsController, configuration,
	)
}
