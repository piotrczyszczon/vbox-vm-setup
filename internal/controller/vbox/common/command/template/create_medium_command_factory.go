package template

import (
	"create-vbox-vm/internal/controller/configuration"
	"create-vbox-vm/internal/controller/vbox/common/command/ds"
	"create-vbox-vm/internal/controller/vbox/common/command/template/argument"
	"create-vbox-vm/internal/controller/vbox/create/scenario/predefined/steps/utils"
)

type CreatMediumCommandFactory struct {
}

func NewCreatMediumCommandFactory() CreatMediumCommandFactory {
	return newCreatMediumCommandFactory()
}

func newCreatMediumCommandFactory() CreatMediumCommandFactory {
	return CreatMediumCommandFactory{}
}

func (factory CreatMediumCommandFactory) CreateCommand(configuration configuration.Configuration) *ds.VBoxManageCommand {
	return &ds.VBoxManageCommand{
		Name: "createmedium",
		Args: []string{
			argument.Arg("filename", utils.GetVmDiskPath(configuration)).AsString(),
			argument.Arg("size", configuration.Create.Predefined.HddSize).AsString(),
			argument.Arg("format", "VDI").AsString(),
		},
	}
}
