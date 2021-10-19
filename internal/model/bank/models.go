package bank

import (
	"fmt"
	"math/big"
)

const currencyDisplayPrecision = 4

type BankAccount struct {
	id          uint64
	UserId      uint64
	IsLegal     bool
	Number      string
	TotalAmount big.Float
	Currency    string
}

func (s *BankAccount) String() string {
	legalStatus := "individual"
	if s.IsLegal {
		legalStatus = "legal"
	}

	return fmt.Sprintf(`%d: userId=%d (%s)
number: %s
value: %s %s
`,
		s.id, s.UserId, legalStatus, s.Number, s.TotalAmount.Text('f', currencyDisplayPrecision), s.Currency,
	)
}

func NewBankAccount(userId uint64, isLegal bool, number string, currency string) BankAccount {
	return BankAccount{
		id:          0,
		UserId:      userId,
		IsLegal:     isLegal,
		Number:      number,
		Currency:    currency,
		TotalAmount: *big.NewFloat(0.0),
	}
}

func CreateWithId(ID uint64, source BankAccount) *BankAccount {
	// TODO подумать про заворачиваниe в id VO на уровне сервиса.

	newAccount := &BankAccount{id: ID}
	newAccount.RefillFromAnother(source)

	return newAccount
}

func (s *BankAccount) RefillFromAnother(account BankAccount) {
	s.UserId = account.UserId
	s.IsLegal = account.IsLegal
	s.Number = account.Number
	s.Currency = account.Currency

	s.SetTotalAmount(account.TotalAmount)
}

func (s *BankAccount) GetId() uint64 {
	return s.id
}

func (s *BankAccount) SetTotalAmount(newTotalAmount big.Float) {
	s.TotalAmount = newTotalAmount
}
