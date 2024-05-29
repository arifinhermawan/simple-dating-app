package server

import "github.com/arifinhermawan/simple-dating-app/internal/usecase/sample"

type UseCase struct {
	Sample *sample.UseCase
}

func NewUseCase() *UseCase {
	return &UseCase{
		Sample: sample.NewUseCase(),
	}
}
