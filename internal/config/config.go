package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var cfg *Config

// GetConfigInstance returns service config
func GetConfigInstance() Config {
	if cfg != nil {
		return *cfg
	}

	return Config{}
}

// Kafka - contains all parameters kafka information.
type Kafka struct {
	Capacity uint64   `yaml:"capacity"`
	Topic    string   `yaml:"topic"`
	GroupID  string   `yaml:"groupId"`
	Brokers  []string `yaml:"brokers"`
}

// Config - contains all configuration parameters in config package.
type Config struct {
	Kafka Kafka `yaml:"kafka"`
}

// ReadConfigYML - read configurations from file and init instance Config.
func ReadConfigYML(filePath string) error {
	if cfg != nil {
		return nil
	}

	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return err
	}

	return nil
}
