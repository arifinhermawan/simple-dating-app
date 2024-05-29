package server

import "github.com/arifinhermawan/simple-dating-app/internal/app/infrastructure/configuration"

type configProvider interface {
	GetConfig() configuration.AppConfig
}

type Infra struct {
	Config configProvider
}

func NewInfra() *Infra {
	return &Infra{
		Config: configuration.NewConfig(),
	}
}
