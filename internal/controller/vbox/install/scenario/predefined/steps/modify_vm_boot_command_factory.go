package steps

import (
	"fmt"
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/common/command/ds"
	"vbox-vm-setup/internal/controller/vbox/common/command/template"
)

type ModifyVmBoot1CommandFactory struct {
	modifyVmCommandFactory template.ModifyVmCommandFactory
	bootId                 string
	bootTarget             string
}

func NewModifyVmBoot1CommandFactory() ModifyVmBoot1CommandFactory {
	return newModifyVmBoot1CommandFactory("1", "dvd", template.NewModifyVmCommandFactory())
}

func NewModifyVmBoot2CommandFactory() ModifyVmBoot1CommandFactory {
	return newModifyVmBoot1CommandFactory("2", "disk", template.NewModifyVmCommandFactory())
}

func NewModifyVmBoot3CommandFactory() ModifyVmBoot1CommandFactory {
	return newModifyVmBoot1CommandFactory("3", "none", template.NewModifyVmCommandFactory())
}

func NewModifyVmBoot4CommandFactory() ModifyVmBoot1CommandFactory {
	return newModifyVmBoot1CommandFactory("4", "none", template.NewModifyVmCommandFactory())
}

func newModifyVmBoot1CommandFactory(bootId string, bootTarget string, modifyVmCommandFactory template.ModifyVmCommandFactory) ModifyVmBoot1CommandFactory {
	return ModifyVmBoot1CommandFactory{bootId: bootId, bootTarget: bootTarget, modifyVmCommandFactory: modifyVmCommandFactory}
}

func (factory ModifyVmBoot1CommandFactory) CreateCommand(configuration configuration.Configuration) *ds.VBoxManageCommand {
	return factory.modifyVmCommandFactory.CreateCommand(
		factory.getParamName(), factory.bootTarget, configuration,
	)
}

func (factory ModifyVmBoot1CommandFactory) getParamName() string {
	return fmt.Sprintf("boot%s", factory.bootId)
}
