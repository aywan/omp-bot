package subdomain

type ServiceInterface interface {
	Describe(ID uint64) (*Subdomain, error)
	List(cursor uint64, limit uint64) ([]Subdomain, error)
	Create(subdomain Subdomain) (uint64, error)
	Update(ID uint64, subdomain Subdomain) error
	Remove(ID uint64) (bool, error)
}
