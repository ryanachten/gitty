package models

import (
	"flag"
	"os"
)

type Arguments struct {
	ConfigurationPath string
	Command           []string
}

// Parses command line arguments
func GetArguments() *Arguments {
	var configPath string
	flag.StringVar(&configPath, "config", "", "Path for gitty configuration file")

	flag.Parse()

	// All trailing args are interpreted the git command to be executed
	trailingArgs := flag.Args()

	if configPath == "" {
		println("Missing -config flag with path to configuration file.\n")
		println("gitty usage:")
		flag.PrintDefaults()
		os.Exit(1)
		return nil
	}

	args := Arguments{
		ConfigurationPath: configPath,
		Command:           trailingArgs,
	}

	return &args
}
