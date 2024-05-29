package server

import "github.com/arifinhermawan/simple-dating-app/internal/service/sample"

type Service struct {
	Sample *sample.Service
}

func NewService() *Service {
	return &Service{
		Sample: sample.NewService(),
	}
}
