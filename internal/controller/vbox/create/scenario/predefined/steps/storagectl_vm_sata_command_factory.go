package steps

import (
	"create-vbox-vm/internal/controller/configuration"
	"create-vbox-vm/internal/controller/vbox/common/command/ds"
	"create-vbox-vm/internal/controller/vbox/common/command/template"
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
