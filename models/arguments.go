package models

import (
	"flag"
	"os"
)

type Arguments struct {
	ConfigurationPath string
}

// Parses command line arguments
func GetArguments() *Arguments {
	var configPath string
	flag.StringVar(&configPath, "config", "", "Path for gitty configuration file")

	flag.Parse()

	if configPath == "" {
		println("Missing -config flag with path to configuration file.\n")
		println("cmdo usage:")
		flag.PrintDefaults()
		os.Exit(1)
		return nil
	}

	args := Arguments{
		ConfigurationPath: configPath,
	}

	return &args
}
