package util

import (
	"bytes"
	"github.com/goccy/go-yaml"
	"os"
	"store/types/cfg"
)

var Config *cfg.Config

func GetConfig(path string) (config *cfg.Config, err error) {
	rawConf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config = new(cfg.Config)
	decoder := yaml.NewDecoder(bytes.NewReader(rawConf))
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}
	Config = config
	return config, nil
}

func DumpConfig(path string, config *cfg.Config) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	if err := yaml.NewEncoder(file).Encode(config); err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}
