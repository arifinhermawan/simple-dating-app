package server

import "github.com/arifinhermawan/simple-dating-app/internal/handler/sample"

type Handler struct {
	Sample *sample.Handler
}

func NewHandler() *Handler {
	return &Handler{
		Sample: sample.NewHandler(),
	}
}
