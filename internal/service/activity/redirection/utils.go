package redirection

import (
	"errors"

	"github.com/ozonmp/omp-bot/internal/service/activity"
)

func validateInput(redirection activity.Redirection, typeOperation string) error {
	if redirection.Title == "" {
		return errors.New("field \"Title\" is empty")
	}

	if redirection.From == "" {
		return errors.New("field \"From\" is empty")
	}

	if redirection.To == "" {
		return errors.New("field \"To\" is empty")
	}

	if typeOperation == "UPDATE" && redirection.Id == 0 {
		return errors.New("field \"id\" is empty")
	}

	return nil
}

func sequence(r *DummyActivityRedirectionService) uint64 {
	r.seqId++
	return r.seqId
}
