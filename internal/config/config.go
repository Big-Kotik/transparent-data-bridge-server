package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type BridgeConfig struct {
	Writer   WriterConfig `yaml:"writerConfig"`
	Endpoint string       `yaml:"endpoint"`
}

type WriterConfig struct {
	ChunkSize int64  `yaml:"chunkSize"` // chunk size in bytes
	BasicDir  string `yaml:"basicDir"`
}

func ConfigFromFile(file string) (*BridgeConfig, error) {
	withArgs := os.ExpandEnv(file)

	var cfg BridgeConfig
	if err := yaml.Unmarshal([]byte(withArgs), &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
