package template

import (
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/common/command/ds"
	"vbox-vm-setup/internal/controller/vbox/common/command/template/argument"
)

type UnattendedInstallCommandFactory struct {
}

func NewUnattendedInstallCommandFactory() UnattendedInstallCommandFactory {
	return newUnattendedInstallCommandFactory()
}

func newUnattendedInstallCommandFactory() UnattendedInstallCommandFactory {
	return UnattendedInstallCommandFactory{}
}

func (factory UnattendedInstallCommandFactory) CreateCommand(configuration configuration.Configuration) *ds.VBoxManageCommand {
	return &ds.VBoxManageCommand{
		Name: "unattended",
		Args: []string{
			argument.Param("install").AsString(),
			argument.Param(configuration.General.Name).AsString(),
			argument.Arg("iso", configuration.General.IsoFilePath).AsString(),
			argument.Arg("full-user-name", configuration.Install.Predefined.FullUserName).AsString(),
			argument.Arg("user", configuration.Install.Predefined.Credentials.User).AsString(),
			argument.Arg("password", configuration.Install.Predefined.Credentials.Password).AsString(),
			argument.Opt("install-additions").AsString(),
			argument.Arg("time-zone", configuration.Install.Predefined.TimeZone).AsString(),
			argument.Arg("post-install-command", configuration.Install.Predefined.PostInstallCommand).AsString(),
		},
	}
}
