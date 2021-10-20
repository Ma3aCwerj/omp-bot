package redirection

import "github.com/ozonmp/omp-bot/internal/service/activity"

type RedirectionService interface {
	Describe(redirectionID uint64) (*activity.Redirection, error)
	List(cursor uint64, limit uint64) ([]activity.Redirection, error)
	Create(activity.Redirection) (uint64, error)
	Update(redirectionID uint64, redirection activity.Redirection) error
	Remove(redirectionID uint64) (bool, error)
}

type DummyActivityRedirectionService struct {
	entities map[uint64]activity.Redirection
	seqId    uint64
}

func NewDummyActivityRedirectionService() *DummyActivityRedirectionService {
	return &DummyActivityRedirectionService{make(map[uint64]activity.Redirection), 0}
}
