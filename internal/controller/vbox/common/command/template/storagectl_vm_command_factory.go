package template

import (
	"create-vbox-vm/internal/controller/vbox/common/command/ds"
	"create-vbox-vm/internal/controller/vbox/common/command/template/argument"
)

type StorageCtlCommandFactory struct {
}

func NewStorageCtlCommandFactory() StorageCtlCommandFactory {
	return newStorageCtlCommandFactory()
}

func newStorageCtlCommandFactory() StorageCtlCommandFactory {
	return StorageCtlCommandFactory{}
}

func (factory StorageCtlCommandFactory) CreateCommand(vmName string, name string, busType string, controller string, portCount string, bootable string, hostiocache string) *ds.VBoxManageCommand {
	return &ds.VBoxManageCommand{
		Name: "storagectl",
		Args: []string{
			argument.Param(vmName).AsString(),
			argument.Arg("name", name).AsString(),
			argument.Arg("add", busType).AsString(),
			argument.Arg("controller", controller).AsString(),
			argument.Arg("portcount", portCount).AsString(),
			argument.Arg("bootable", bootable).AsString(),
			argument.Arg("hostiocache", hostiocache).AsString(),
		},
	}
}
