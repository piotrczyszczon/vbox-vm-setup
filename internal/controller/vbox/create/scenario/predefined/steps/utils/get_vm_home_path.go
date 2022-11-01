package utils

import (
	"create-vbox-vm/internal/controller/configuration"
	"fmt"
)

func GetVmHomePath(configuration configuration.Configuration) string {
	return fmt.Sprintf("%s/%s", configuration.Create.Predefined.VmsHome, configuration.General.Name)
}
