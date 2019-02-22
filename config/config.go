package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var defaultConfig = Config{
	Name:    "3pxls",
	Root:    "./",
	Storage: "./files",
	Port:    ":8080",
	Secret:  "", // empty secret will not allow new images to be uploaded.
	GAID:    "",
}

// Config contains all the configuration options.
type Config struct {
	Name    string `yaml:"name"`
	Root    string `yaml:"root"`
	Storage string `yaml:"storage"`
	Port    string `yaml:"port"`
	Secret  string `yaml:"secret"`
	GAID    string `yaml:"GAID"`
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

func LoadFromEnv() (*Config, error) {
	cfg := defaultConfig

	envKeys := map[string]*string{
		"PIX_NAME":    &cfg.Name,
		"PIX_ROOT":    &cfg.Root,
		"PIX_STORAGE": &cfg.Storage,
		"PIX_PORT":    &cfg.Port,
		"PIX_SECRET":  &cfg.Secret,
		"PIX_GAID":    &cfg.GAID,
	}

	for envvar, target := range envKeys {
		if value, ok := os.LookupEnv(envvar); ok {
			*target = value
		}
	}
	return &cfg, nil
}
