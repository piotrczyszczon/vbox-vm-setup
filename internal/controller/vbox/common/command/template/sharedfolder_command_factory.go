package template

import (
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/common/command/ds"
	"vbox-vm-setup/internal/controller/vbox/common/command/template/argument"
)

type SharedfolderAddCommandFactory struct {
}

func NewSharedfolderAddCommandFactory() SharedfolderAddCommandFactory {
	return newSharedfolderAddCommandFactory()
}

func newSharedfolderAddCommandFactory() SharedfolderAddCommandFactory {
	return SharedfolderAddCommandFactory{}
}

func (provider SharedfolderAddCommandFactory) CreateCommand(sharedFolderName string, hostPath string, autoMountPoint string, configuration configuration.Configuration) *ds.VBoxManageCommand {
	return &ds.VBoxManageCommand{
		Name: "sharedfolder",
		Args: []string{
			argument.Param("add").AsString(),
			argument.Param(configuration.General.Name).AsString(),
			argument.Arg("name", sharedFolderName).AsString(),
			argument.Arg("hostpath", hostPath).AsString(),
			argument.Arg("auto-mount-point", autoMountPoint).AsString(),
			argument.Opt("automount").AsString(),
		},
	}
}
