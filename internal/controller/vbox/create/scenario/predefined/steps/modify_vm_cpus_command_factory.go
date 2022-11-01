package steps

import (
	"create-vbox-vm/internal/controller/configuration"
	"create-vbox-vm/internal/controller/vbox/common/command/ds"
	"create-vbox-vm/internal/controller/vbox/common/command/template"
)

type ModifyVmCpusCommandFactory struct {
	modifyVmCommandFactory template.ModifyVmCommandFactory
}

func NewModifyVmCpusCommandFactory() ModifyVmCpusCommandFactory {
	return newModifyVmCpusCommandFactory(template.NewModifyVmCommandFactory())
}

func newModifyVmCpusCommandFactory(modifyVmCommandFactory template.ModifyVmCommandFactory) ModifyVmCpusCommandFactory {
	return ModifyVmCpusCommandFactory{modifyVmCommandFactory: modifyVmCommandFactory}
}

func (factory ModifyVmCpusCommandFactory) CreateCommand(configuration configuration.Configuration) *ds.VBoxManageCommand {
	return factory.modifyVmCommandFactory.CreateCommand(
		"cpus", configuration.Create.Predefined.Cpus, configuration,
	)
}
