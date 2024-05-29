package configuration

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func (c Configuration) GetConfig() AppConfig {
	c.doLoadConfigOnce.Do(func() {
		cfg, err := c.LoadConfig()
		if err != nil {
			log.Fatalf("[GetConfig] c.LoadConfig() got error: %+v\n", err)
		}

		c.Config = cfg
	})

	return c.Config
}

func (c Configuration) LoadConfig() (AppConfig, error) {
	data, err := os.ReadFile("files/config.development.yaml")
	if err != nil {
		log.Printf("[LoadConfig] os.ReadFile() got error: %+v\n", err)
		return AppConfig{}, err
	}

	var config AppConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Printf("[LoadConfig] yaml.Unmarshal() got error: %+v\n", err)
		return AppConfig{}, err
	}

	return config, nil
}
