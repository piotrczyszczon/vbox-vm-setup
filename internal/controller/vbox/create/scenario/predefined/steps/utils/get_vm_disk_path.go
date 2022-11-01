package utils

import (
	"fmt"
	"vbox-vm-setup/internal/controller/configuration"
)

func GetVmDiskPath(configuration configuration.Configuration) string {
	return fmt.Sprintf("%s/%s.vdi", GetVmHomePath(configuration), configuration.General.Name)
}
