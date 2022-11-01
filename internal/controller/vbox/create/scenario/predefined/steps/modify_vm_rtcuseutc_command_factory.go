package steps

import (
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/common/command/ds"
	"vbox-vm-setup/internal/controller/vbox/common/command/template"
)

type ModifyVmRtcUseUtcCommandFactory struct {
	modifyVmCommandFactory template.ModifyVmCommandFactory
}

func NewModifyVmRtcUseUtcCommandFactory() ModifyVmRtcUseUtcCommandFactory {
	return newModifyVmRtcUseUtcCommandFactory(template.NewModifyVmCommandFactory())
}

func newModifyVmRtcUseUtcCommandFactory(modifyVmCommandFactory template.ModifyVmCommandFactory) ModifyVmRtcUseUtcCommandFactory {
	return ModifyVmRtcUseUtcCommandFactory{modifyVmCommandFactory: modifyVmCommandFactory}
}

func (factory ModifyVmRtcUseUtcCommandFactory) CreateCommand(configuration configuration.Configuration) *ds.VBoxManageCommand {
	return factory.modifyVmCommandFactory.CreateCommand(
		"rtcuseutc", configuration.Create.Predefined.RtcUseUtc, configuration,
	)
}
