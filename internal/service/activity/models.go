package activity

import (
	"fmt"
)

type Redirection struct {
	Id          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	From        string `json:"from"`
	To          string `json:"to"`
}

func (s *Redirection) String() string {
	return fmt.Sprintf("%d - %s\ndescription:\n%s\nFrom - %s\nTo - %s\n", s.Id, s.Title, s.Description, s.From, s.To)
}
