package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Token string `toml:"token"`
}

func (cfg *Config) Initialize() error {
	dir := filepath.Join(os.Getenv("HOME"), ".config", "nature-remo")

	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("cannot create directoty: %v", err)
	}

	file := filepath.Join(dir, "config.toml")

	_, err := os.Stat(file)
	if err == nil {
		_, err = toml.DecodeFile(file, cfg)
		if err != nil {
			return err
		}

		return nil
	}

	if !os.IsNotExist(err) {
		return err
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	cfg.Token = "<YOUR ACCESS TOKEN>"

	return toml.NewEncoder(f).Encode(cfg)
}
