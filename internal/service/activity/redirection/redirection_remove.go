package redirection

import "fmt"

func (s *DummyActivityRedirectionService) Remove(redirectionID uint64) (bool, error) {
	_, ok := s.entities[redirectionID]
	if !ok {
		return false, fmt.Errorf("redirection not found %d", redirectionID)
	}

	delete(s.entities, redirectionID)

	return true, nil
}
