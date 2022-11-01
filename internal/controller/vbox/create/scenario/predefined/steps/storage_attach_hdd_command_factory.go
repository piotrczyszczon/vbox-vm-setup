package steps

import (
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/common/command/ds"
	"vbox-vm-setup/internal/controller/vbox/common/command/template"
	"vbox-vm-setup/internal/controller/vbox/create/scenario/predefined/steps/utils"
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
