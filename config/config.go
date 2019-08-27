package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	RPCUrl      string `yaml:"rpc"`
	HTTPAddress string `yaml:"address"`
}

func (c *Config) Load(name string) error {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(buf, c)
}
