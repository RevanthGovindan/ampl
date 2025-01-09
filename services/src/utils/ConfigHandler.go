package utils

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("%s is a directory, not a normal file", path)
	}
	return nil
}

func parseConfigFile[S any](configPath string, config *S) error {
	buf, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, config)
	if err != nil {
		return err
	}
	return nil
}

func InitializeConfigs[S any](configRef *S) error {
	var configPath string
	flag.StringVar(&configPath, "config", "config.yaml", "path to config file")
	flag.Parse()
	if err := ValidateConfigPath(configPath); err != nil {
		return err
	}
	return parseConfigFile(configPath, configRef)
}
