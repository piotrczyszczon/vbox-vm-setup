package steps

import (
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/common/command/ds"
	"vbox-vm-setup/internal/controller/vbox/common/command/template"
)

type ModifyVmClipboardModeCommandFactory struct {
	modifyVmCommandFactory template.ModifyVmCommandFactory
}

func NewModifyVmClipboardModeCommandFactory() ModifyVmClipboardModeCommandFactory {
	return newModifyVmClipboardModeCommandFactory(template.NewModifyVmCommandFactory())
}

func newModifyVmClipboardModeCommandFactory(modifyVmCommandFactory template.ModifyVmCommandFactory) ModifyVmClipboardModeCommandFactory {
	return ModifyVmClipboardModeCommandFactory{modifyVmCommandFactory: modifyVmCommandFactory}
}

func (factory ModifyVmClipboardModeCommandFactory) CreateCommand(configuration configuration.Configuration) *ds.VBoxManageCommand {
	return factory.modifyVmCommandFactory.CreateCommand(
		"clipboard", configuration.Create.Predefined.ClipboardMode, configuration,
	)
}
