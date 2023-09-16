package config

import (
	"errors"
	"io"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type HttpServer struct {
	Address string        `yaml:"address"`
	Timeout time.Duration `yaml:"time_out"`
}

type Config struct {
	Env         string `yaml:"env"`
	StoragePath string `yaml:"storage_path"`
	HttpServer  `yaml:"http_server"`
}

func parseYamlToConfig(f *os.File) (*Config, error) {
	var cfg Config
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, errors.New("Failed to read file")
	}

	yamlErr := yaml.Unmarshal(data, &cfg)
	if yamlErr != nil {
		return nil, yamlErr
	}

	return &cfg, nil
}

func Load() (*Config, error) {
	cfgPath := os.Getenv("CONFIG_PATH")

	if cfgPath == "" {
		return nil, errors.New("Can't find config file, set up CONFIG_PATH env variable")
		//panic("Can't find config file, set up CONFIG_PATH env variable")

	}

	cfgFile, err := os.Open(cfgPath)
	if err != nil {
		return nil, errors.New("Failed to open config file")
		//panic("Failed to open config file")
	}
	defer cfgFile.Close()

	cfg, err := parseYamlToConfig(cfgFile)
	if err != nil {
		return nil, errors.New("Failed to parse config file")
		//panic("Failed to parse config file")
	}

	return cfg, nil

}
