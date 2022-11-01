package steps

import (
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/common/command/ds"
	"vbox-vm-setup/internal/controller/vbox/common/command/template"
)

type StorageAttachDvdDriveCommandFactory struct {
	storageAttachCommandFactory template.StorageAttachCommandFactory
}

func NewStorageAttachDvdDriveCommandFactory() StorageAttachDvdDriveCommandFactory {
	return newStorageAttachDvdDriveCommandFactory(template.NewStorageAttachCommandFactory())
}

func newStorageAttachDvdDriveCommandFactory(storageAttachCommandFactory template.StorageAttachCommandFactory) StorageAttachDvdDriveCommandFactory {
	return StorageAttachDvdDriveCommandFactory{storageAttachCommandFactory: storageAttachCommandFactory}
}

func (factory StorageAttachDvdDriveCommandFactory) CreateCommand(configuration configuration.Configuration) *ds.VBoxManageCommand {
	return factory.storageAttachCommandFactory.CreateCommand(
		configuration.General.Name,
		"IDE",
		"dvddrive",
		"1",
		"0",
		configuration.General.IsoFilePath,
	)
}
