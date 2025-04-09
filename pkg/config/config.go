package config

import "github.com/BurntSushi/toml"

const (
	version string = "0.1"
)

type Config struct {
	ConfFile string
	Addr     string `toml:"addr"`
	Port     int    `toml:"port"`
	Database string `toml:"database"`
	Version  string
}

func (c *Config) GetConfig() {
	_, err := toml.DecodeFile(c.ConfFile, &c)
	if err != nil {
		panic(err)
	}
	c.Version = version
}
