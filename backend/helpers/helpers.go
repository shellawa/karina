package helpers

import "context"

type Service struct {
	ctx context.Context
}

func (s *Service) Initialize(ctx context.Context) {
	s.ctx = ctx
}
