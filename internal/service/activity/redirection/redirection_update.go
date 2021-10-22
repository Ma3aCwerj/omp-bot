package redirection

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/service/activity"
)

func (s *DummyActivityRedirectionService) Update(redirectionID uint64, redirection activity.Redirection) error {
	if err := validateInput(redirection, "UPDATE"); err != nil {
		return err
	}

	entity, ok := s.entities[redirectionID]
	if !ok {
		return fmt.Errorf("redirecction not found %d", redirectionID)
	}

	entity.Title = redirection.Title
	entity.Description = redirection.Description
	entity.From = redirection.From
	entity.To = redirection.To
	s.entities[redirectionID] = entity
	return nil
}
