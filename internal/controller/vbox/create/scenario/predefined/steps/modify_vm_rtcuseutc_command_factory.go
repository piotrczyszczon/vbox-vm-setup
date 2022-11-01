package steps

import (
	"create-vbox-vm/internal/controller/configuration"
	"create-vbox-vm/internal/controller/vbox/common/command/ds"
	"create-vbox-vm/internal/controller/vbox/common/command/template"
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
