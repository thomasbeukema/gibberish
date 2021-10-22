package config

import (
	"fmt"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Struct which holds the config options
// for the YAML file
type Configuration struct {
	ApiKey string `yaml:"ApiKey"`
}

func LoadConfig() *Configuration {
	// Create new Configuration struct
	conf := Configuration{}

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Read the config.yaml file
	confFile, err := os.ReadFile(fmt.Sprintf("%s%s", cwd, "/config.yaml"))
	if err != nil {
		log.Fatal(err)
	}

	// Parse the YAML
	err = yaml.Unmarshal(confFile, &conf)
	if err != nil {
		log.Fatal(err)
	}

	return &conf
}
