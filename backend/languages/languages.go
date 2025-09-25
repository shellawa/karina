package languages

import (
	"context"
	"karina/backend/languages/python"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Runner interface {
	Name() string
	Run(sourcePath string) string
}

type Service struct {
	ctx     context.Context
	runners map[string]Runner
}

func (s *Service) Initialize(ctx context.Context) {
	s.ctx = ctx
	s.runners = map[string]Runner{
		"python": &python.Runner{},
	}
}

func (s *Service) RunSample() {
	v := s.runners["python"].Run("dsds")
	runtime.EventsEmit(s.ctx, "run:python:sample", v)
}
