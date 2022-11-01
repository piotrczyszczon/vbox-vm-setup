package steps

import (
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/common/command/ds"
	"vbox-vm-setup/internal/controller/vbox/common/command/template"
)

type StorageCtlSataCommandFactory struct {
	storageCtlCommandFactory template.StorageCtlCommandFactory
}

func NewStorageCtlSataCommandFactory() StorageCtlSataCommandFactory {
	return newStorageCtlSataCommandFactory(template.NewStorageCtlCommandFactory())
}

func newStorageCtlSataCommandFactory(storageCtlCommandFactory template.StorageCtlCommandFactory) StorageCtlSataCommandFactory {
	return StorageCtlSataCommandFactory{storageCtlCommandFactory: storageCtlCommandFactory}
}

func (factory StorageCtlSataCommandFactory) CreateCommand(configuration configuration.Configuration) *ds.VBoxManageCommand {
	return factory.storageCtlCommandFactory.CreateCommand(
		configuration.General.Name,
		"SATA",
		"sata",
		"IntelAhci",
		"30",
		"on",
		"off",
	)
}
