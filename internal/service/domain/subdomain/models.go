package subdomain

import "fmt"

type Subdomain struct {
	id    uint64
	Title string
}

func (s *Subdomain) String() string {
	return fmt.Sprintf("%d: %s", s.id, s.Title)
}

func NewSubdomain(title string) Subdomain {
	return Subdomain{
		id:    0,
		Title: title,
	}
}

func recreateSubdomain(ID uint64, s Subdomain) *Subdomain {
	// TODO подумать про заворачиваниe в ID на уровне сервиса.
	return &Subdomain{
		id:    ID,
		Title: s.Title,
	}
}

func (s *Subdomain) fillFrom(subdomain Subdomain) {
	s.Title = subdomain.Title
}

func (s *Subdomain) GetId() uint64 {
	return s.id
}
