package utils

import (
	"create-vbox-vm/internal/controller/configuration"
	"fmt"
)

func GetVmDiskPath(configuration configuration.Configuration) string {
	return fmt.Sprintf("%s/%s.vdi", GetVmHomePath(configuration), configuration.General.Name)
}
