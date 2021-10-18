package bank_account

type ServiceInterface interface {
	Describe(ID uint64) (*BankAccount, error)
	List(cursor uint64, limit uint64) ([]BankAccount, error)
	Create(subdomain BankAccount) (uint64, error)
	Update(ID uint64, subdomain BankAccount) error
	Remove(ID uint64) (bool, error)
}
