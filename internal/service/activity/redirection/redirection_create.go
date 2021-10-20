package redirection

import "github.com/ozonmp/omp-bot/internal/service/activity"

func (s *DummyActivityRedirectionService) Create(redirection activity.Redirection) (uint64, error) {

	if err := validateInput(redirection, "CREATE"); err != nil {
		return 0, err
	}

	redirection.Id = sequence(s)

	s.entities[redirection.Id] = redirection
	return redirection.Id, nil
}
