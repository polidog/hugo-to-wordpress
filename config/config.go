package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	WordPress WordPressConfig `yaml:"wordpress"`
	Hugo      HugoConfig      `yaml:"hugo"`
}

type WordPressConfig struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type HugoConfig struct {
	ContentPath string `yaml:"content_path"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
