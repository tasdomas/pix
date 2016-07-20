package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var defaultConfig = Config{
	Name:    "3pxls",
	Root:    "./",
	Storage: "./files",
	Port:    ":8080",
	Secret:  "", // empty secret will not allow new images to be uploaded.
}

// Config contains all the configuration options.
type Config struct {
	Name    string `yaml:"name"`
	Root    string `yaml:"root"`
	Storage string `yaml:"storage"`
	Port    string `yaml:"port"`
	Secret  string `yaml:"secret"`
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
