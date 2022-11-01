package steps

import (
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/common/command/ds"
	"vbox-vm-setup/internal/controller/vbox/common/command/template"
)

type ModifyVmNatRuleCommandFactory struct {
	modifyVmCommandFactory template.ModifyVmCommandFactory
}

func NewModifyVmNatRuleCommandFactory() ModifyVmNatRuleCommandFactory {
	return newModifyVmNatRuleCommandFactory(template.NewModifyVmCommandFactory())
}

func newModifyVmNatRuleCommandFactory(modifyVmCommandFactory template.ModifyVmCommandFactory) ModifyVmNatRuleCommandFactory {
	return ModifyVmNatRuleCommandFactory{modifyVmCommandFactory: modifyVmCommandFactory}
}

func (factory ModifyVmNatRuleCommandFactory) CreateCommand(configuration configuration.Configuration) *ds.VBoxManageCommand {
	if len(configuration.Create.Predefined.NatRule) > 0 {
		return factory.modifyVmCommandFactory.CreateCommand(
			"natpf1", configuration.Create.Predefined.NatRule, configuration,
		)
	} else {
		return nil
	}
}
