package steps

import (
	"create-vbox-vm/internal/controller/configuration"
	"create-vbox-vm/internal/controller/vbox/common/command/ds"
	"create-vbox-vm/internal/controller/vbox/common/command/template"
)

type StorageCtlIdeCommandFactory struct {
	storageCtlCommandFactory template.StorageCtlCommandFactory
}

func NewStorageCtlIdeCommandFactory() StorageCtlIdeCommandFactory {
	return newStorageCtlIdeCommandFactory(template.NewStorageCtlCommandFactory())
}

func newStorageCtlIdeCommandFactory(storageCtlCommandFactory template.StorageCtlCommandFactory) StorageCtlIdeCommandFactory {
	return StorageCtlIdeCommandFactory{storageCtlCommandFactory: storageCtlCommandFactory}
}

func (factory StorageCtlIdeCommandFactory) CreateCommand(configuration configuration.Configuration) *ds.VBoxManageCommand {
	return factory.storageCtlCommandFactory.CreateCommand(
		configuration.General.Name,
		"IDE",
		"ide",
		"PIIX4",
		"2",
		"on",
		"on",
	)
}
