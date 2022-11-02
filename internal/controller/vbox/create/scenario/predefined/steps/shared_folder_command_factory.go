package steps

import (
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/common/command/ds"
	"vbox-vm-setup/internal/controller/vbox/common/command/template"
)

type SharedfolderCommandFactory struct {
	sharedfolderAddCommandFactory template.SharedfolderAddCommandFactory
}

func NewSharedfolderCommandFactory() SharedfolderCommandFactory {
	return newSharedfolderCommandFactory(template.NewSharedfolderAddCommandFactory())
}

func newSharedfolderCommandFactory(sharedfolderAddCommandFactory template.SharedfolderAddCommandFactory) SharedfolderCommandFactory {
	return SharedfolderCommandFactory{sharedfolderAddCommandFactory: sharedfolderAddCommandFactory}
}

func (factory SharedfolderCommandFactory) CreateCommand(configuration configuration.Configuration) *ds.VBoxManageCommand {
	if len(configuration.Create.Predefined.SharedFolder.Name) > 0 {
		sharedFolderName := configuration.General.Name + configuration.Create.Predefined.SharedFolder.Name

		return factory.sharedfolderAddCommandFactory.CreateCommand(
			sharedFolderName,
			configuration.Create.Predefined.SharedFolder.HostPath,
			configuration.Create.Predefined.SharedFolder.AutoMountPoint,
			configuration,
		)
	} else {
		return nil
	}
}
