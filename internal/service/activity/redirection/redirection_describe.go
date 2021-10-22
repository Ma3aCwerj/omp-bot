package redirection

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/service/activity"
)

func (s *DummyActivityRedirectionService) Describe(redirectionID uint64) (*activity.Redirection, error) {
	entitiy, ok := s.entities[redirectionID]
	if !ok {
		return nil, fmt.Errorf("redirection not found %d", redirectionID)
	}

	return &entitiy, nil
}
