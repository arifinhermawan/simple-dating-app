package server

import (
	"io"

	"github.com/arifinhermawan/simple-dating-app/internal/app/infrastructure/configuration"
	"github.com/arifinhermawan/simple-dating-app/internal/app/infrastructure/golang"
)

type configProvider interface {
	GetConfig() *configuration.AppConfig
}

type golangProvider interface {
	JsonMarshal(input interface{}) ([]byte, error)
	JsonUnmarshal(input []byte, dest interface{}) error
	ReadAll(input io.Reader) ([]byte, error)
}

type Infra struct {
	Config configProvider
	Golang golangProvider
}

func NewInfra() *Infra {
	return &Infra{
		Config: configuration.NewConfig(),
		Golang: golang.NewGolang(),
	}
}

func (i *Infra) GetConfig() *configuration.AppConfig {
	return i.Config.GetConfig()
}

func (i *Infra) JsonMarshal(input interface{}) ([]byte, error) {
	return i.Golang.JsonMarshal(input)
}

func (i *Infra) JsonUnmarshal(input []byte, dest interface{}) error {
	return i.Golang.JsonUnmarshal(input, dest)
}

func (i *Infra) ReadAll(input io.Reader) ([]byte, error) {
	return i.Golang.ReadAll(input)
}
