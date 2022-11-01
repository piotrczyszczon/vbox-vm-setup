package cli

import (
	"create-vbox-vm/internal/controller/commons"
	"fmt"
	"os"
)

func cliPrintUsage() {
	applicationName := commons.GetApplicationName()

	fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", applicationName)
	fmt.Fprintf(os.Stderr, "\t -%s <file_path> - path to the configuration file [.yml]\n", OPTION_CONFIG_PATH)
	fmt.Fprintf(os.Stderr, "\t -%s - enable debug messages\n", OPTION_DEBUG)
}
