package steps

import (
	"create-vbox-vm/internal/controller/configuration"
	"create-vbox-vm/internal/controller/vbox/common/command/ds"
	"create-vbox-vm/internal/controller/vbox/common/command/template"
	"create-vbox-vm/internal/controller/vbox/create/scenario/predefined/steps/utils"
)

type StorageAttachHddCommandFactory struct {
	storageAttachCommandFactory template.StorageAttachCommandFactory
}

func NewStorageAttachHddCommandFactory() StorageAttachHddCommandFactory {
	return newStorageAttachHddCommandFactory(template.NewStorageAttachCommandFactory())
}

func newStorageAttachHddCommandFactory(storageAttachCommandFactory template.StorageAttachCommandFactory) StorageAttachHddCommandFactory {
	return StorageAttachHddCommandFactory{storageAttachCommandFactory: storageAttachCommandFactory}
}

func (factory StorageAttachHddCommandFactory) CreateCommand(configuration configuration.Configuration) *ds.VBoxManageCommand {
	return factory.storageAttachCommandFactory.CreateCommand(
		configuration.General.Name,
		"SATA",
		"hdd",
		"0",
		"0",
		utils.GetVmDiskPath(configuration),
	)
}
