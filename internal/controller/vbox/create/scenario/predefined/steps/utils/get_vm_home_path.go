package utils

import (
	"fmt"
	"vbox-vm-setup/internal/controller/configuration"
)

func GetVmHomePath(configuration configuration.Configuration) string {
	return fmt.Sprintf("%s/%s", configuration.Create.Predefined.VmsHome, configuration.General.Name)
}
