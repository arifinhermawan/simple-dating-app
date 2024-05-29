package configuration

import "sync"

type Configuration struct {
	Config AppConfig

	doLoadConfigOnce *sync.Once
}

func NewConfig() *Configuration {
	return &Configuration{
		doLoadConfigOnce: new(sync.Once),
	}
}
