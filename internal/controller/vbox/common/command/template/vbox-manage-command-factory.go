package template

import (
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/common/command/ds"
)

type VBoxManageCommandFactory interface {
	CreateCommand(configuration configuration.Configuration) *ds.VBoxManageCommand
}
