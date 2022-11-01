package template

import (
	"vbox-vm-setup/internal/controller/vbox/common/command/ds"
	"vbox-vm-setup/internal/controller/vbox/common/command/template/argument"
)

type StorageAttachCommandFactory struct {
}

func NewStorageAttachCommandFactory() StorageAttachCommandFactory {
	return newStorageAttachCommandFactory()
}

func newStorageAttachCommandFactory() StorageAttachCommandFactory {
	return StorageAttachCommandFactory{}
}

func (factory StorageAttachCommandFactory) CreateCommand(vmName string, storagectl string, storageType string, port string, device string, medium string) *ds.VBoxManageCommand {
	return &ds.VBoxManageCommand{
		Name: "storageattach",
		Args: []string{
			argument.Param(vmName).AsString(),
			argument.Arg("storagectl", storagectl).AsString(),
			argument.Arg("type", storageType).AsString(),
			argument.Arg("port", port).AsString(),
			argument.Arg("device", device).AsString(),
			argument.Arg("medium", medium).AsString(),
		},
	}
}
