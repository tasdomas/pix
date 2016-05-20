package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var defaultConfig = Config{
	Root:    "./",
	Storage: "./files",
	Port:    ":8080",
}

// Config contains all the configuration options.
type Config struct {
	Root    string `yaml:"root"`
	Storage string `yaml:"storage"`
	Port    string `yaml:"port"`
}

// Load loads the configuration file.
func Load(file string) (*Config, error) {
	cfg := defaultConfig
	if file == "" {
		return &cfg, nil
	}
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(raw, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
