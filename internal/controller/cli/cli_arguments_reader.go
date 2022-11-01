package cli

import (
	"flag"
	"fmt"
	"os"
)

const EMPTY = ""

const (
	OPTION_CONFIG_PATH string = "config-path"
	OPTION_DEBUG       string = "debug"
)

type PrintUsage func()

type CliArgumentsReader struct {
	printUsage PrintUsage
}

func NewCliArgumentsReader() CliArgumentsReader {
	return newCliArgumentsReader(cliPrintUsage)
}

func newCliArgumentsReader(printUsage PrintUsage) CliArgumentsReader {
	return CliArgumentsReader{printUsage: printUsage}
}

func (cliArgumentsReader CliArgumentsReader) GetCliArguments() (*CliArguments, error) {
	optionsFlagSet := flag.NewFlagSet(EMPTY, flag.ContinueOnError)
	optionsFlagSet.Usage = cliArgumentsReader.printUsage

	configPathFlag := optionsFlagSet.String(OPTION_CONFIG_PATH, EMPTY, EMPTY)
	debugFlag := optionsFlagSet.Bool(OPTION_DEBUG, false, EMPTY)

	if len(os.Args) < 1 {
		cliArgumentsReader.printUsage()
		return nil, fmt.Errorf("Not enough arguments")
	}

	parseError := optionsFlagSet.Parse(os.Args[1:])
	if parseError != nil && parseError != flag.ErrHelp {
		return nil, parseError
	}

	if !optionsFlagSet.Parsed() {
		cliArgumentsReader.printUsage()
		return nil, fmt.Errorf("Invalid options provided")
	}

	return &CliArguments{
		ConfigPath: *configPathFlag,
		Debug:      *debugFlag,
	}, nil
}
