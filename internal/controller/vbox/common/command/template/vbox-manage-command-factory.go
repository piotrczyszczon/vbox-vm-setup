package template

import (
	"create-vbox-vm/internal/controller/configuration"
	"create-vbox-vm/internal/controller/vbox/common/command/ds"
)

type VBoxManageCommandFactory interface {
	CreateCommand(configuration configuration.Configuration) *ds.VBoxManageCommand
}
