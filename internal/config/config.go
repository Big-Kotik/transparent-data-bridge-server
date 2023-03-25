package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type BridgeConfig struct {
	Writer        WriterConfig `yaml:"writerConfig"`
	ProxyEndpoint string       `yaml:"proxyEndpoint"`
	Workers       int          `yaml:"workers"`
	Id            int32        `yaml:"id"`
}

type WriterConfig struct {
	BasicDir string `yaml:"basicDir"`
}

func ConfigFromFile(file string) (*BridgeConfig, error) {
	withArgs := os.ExpandEnv(file)

	var cfg BridgeConfig
	if err := yaml.Unmarshal([]byte(withArgs), &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
