package redirection

import (
	"errors"

	"github.com/ozonmp/omp-bot/internal/service/activity"
)

func (s *DummyActivityRedirectionService) List(cursor uint64, limit uint64) ([]activity.Redirection, error) {
	if len(s.entities) == 0 {
		return []activity.Redirection{}, nil
	}

	if cursor > s.seqId {
		return nil, errors.New("redirection not found")
	}

	result := make([]activity.Redirection, 0, limit)

	for i := cursor; i <= s.seqId && limit > 0; i++ {
		if k, ok := s.entities[i]; ok {
			result = append(result, k)
			limit--
		}
	}

	return result, nil
}
