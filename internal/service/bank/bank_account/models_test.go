package bank_account

import (
	"math/big"
	"math/rand"
	"testing"
)

func TestNewBankAccount(t *testing.T) {

	userId := rand.Uint64()
	isLegal := rand.Uint64()%2 == 0

	account := NewBankAccount(userId, isLegal, "0000001", "RUB")

	if account.GetId() != 0 {
		t.Errorf("await id=0, got id=%d", account.GetId())
	}
	if account.IsLegal != isLegal {
		t.Errorf("await isLegal=%t, got isLegal=%t", isLegal, account.IsLegal)
	}
	zeroBigFloat := big.NewFloat(0.0)
	if account.TotalAmount.Cmp(zeroBigFloat) != 0 {
		t.Errorf("await totalAmout=%v, got totalAmount=%v", zeroBigFloat, account.TotalAmount)
	}
}

func TestRecreateBankAccount(t *testing.T) {

	TotalAmount := big.NewFloat(1000.0)
	baseAccount := NewBankAccount(1, true, "0000001", "RUB")
	baseAccount.SetTotalAmount(*TotalAmount)

	id := rand.Uint64()

	account := recreateBankAccount(id, baseAccount)

	if &baseAccount == account {
		t.Error("struct has same addresses")
	}
	if account.GetId() != id {
		t.Errorf("await id=%d, got id=%d", id, account.GetId())
	}
	if account.IsLegal != true {
		t.Errorf("await isLegal=%t, got isLegal=%t", true, account.IsLegal)
	}
	if account.TotalAmount.Cmp(TotalAmount) != 0 {
		t.Errorf("await totalAmout=%v, got totalAmount=%v", TotalAmount, account.TotalAmount)
	}
}
