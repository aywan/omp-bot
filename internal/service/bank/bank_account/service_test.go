package bank_account

import "testing"

func TestEmptyService(t *testing.T) {
	service := NewService()

	items, err := service.List(0, 10)
	if err == nil {
		t.Error("await error, got nothing")
	}
	if len(items) > 0 {
		t.Error("empty service return value")
	}
}

func TestSimpleCRUD(t *testing.T) {
	service := NewService()

	aNumber := "001"
	account := newTestBankAccount(aNumber)

	aId, err := service.Create(*account)
	if err != nil {
		t.Error(err)
	}

	item, err := service.Describe(aId)
	if err != nil {
		t.Error(err)
	}
	if item.Number != aNumber {
		t.Errorf("wrong model title returns %s != %s", item.Number, aNumber)
	}

	bNumber := "002"
	account.Number = bNumber

	err = service.Update(aId, *account)
	if err != nil {
		t.Error(err)
	}

	item, err = service.Describe(aId)
	if err != nil {
		t.Error(err)
	}
	if item.Number != bNumber {
		t.Errorf("wrong model title returns %s != %s", item.Number, bNumber)
	}

	state, err := service.Remove(aId)
	if err != nil {
		t.Error(err)
	}
	if !state {
		t.Errorf("error on delete id=%d", aId)
	}

	state, err = service.Remove(aId)
	if err == nil {
		t.Error("double delete without error")
	}
	if state {
		t.Error("double delete returns true")
	}

	item, err = service.Describe(aId)
	if err == nil {
		t.Error("await error, nil returns")
	}
	if nil != item {
		t.Error("await nil, model returns")
	}
}

func TestServiceList(t *testing.T) {
	service := NewService()

	numbers := []string{
		"001",
		"002",
		"003",
		"004",
		"005",
	}

	numberIds := make(map[string]uint64)

	for _, title := range numbers {
		account := newTestBankAccount(title)
		id, err := service.Create(*account)
		if err != nil {
			t.Errorf("error when create new: %v", err)
		}
		numberIds[title] = id
	}

	items, err := service.List(0, 2)
	if err != nil {
		t.Errorf("error when get list: %v", err)
	}
	if len(items) != 2 {
		t.Errorf("await list len 2, got %d items", len(items))
	}
	if items[0].Number != "001" && items[1].Number != "002" {
		t.Errorf("await [001,002] numbers, got [%s,%s]", items[0].Number, items[1].Number)
	}

	items, err = service.List(items[1].GetId(), 3)

	if err != nil {
		t.Errorf("error when get list: %v", err)
	}
	if len(items) != 3 {
		t.Errorf("await list len 3, got %d items", len(items))
	}
	if items[0].Number != "003" && items[1].Number != "004" && items[2].Number != "005" {
		t.Errorf("await [003,004,005] numbers, got [%s,%s,%s]", items[0].Number, items[1].Number, items[2].Number)
	}

	items, err = service.List(items[2].GetId(), 10)
	if err == nil {
		t.Error("await error, got nothing")
	}

	toRemoveTitles := []string{"003", "005", "001"}
	for _, title := range toRemoveTitles {
		id := numberIds[title]

		result, err := service.Remove(id)
		if err != nil || !result {
			t.Errorf("error when remove id=%d title=%s", id, title)
		}
	}

	items, err = service.List(0, 100)
	if err != nil {
		t.Errorf("error when get list: %v", err)
	}
	if len(items) != 2 {
		t.Errorf("await list len 2, got %d items", len(items))
	}
	if items[0].Number != "002" && items[1].Number != "004" {
		t.Errorf("await [002,004] numbers, got [%s,%s]", items[0].Number, items[1].Number)
	}

	_, err = service.Create(*newTestBankAccount("006"))
	if err != nil {
		t.Errorf("error when create new: %v", err)
	}
	items, err = service.List(items[1].GetId(), 100)
	if len(items) != 1 {
		t.Errorf("await list len 1, got %d items", len(items))
	}
	if items[0].Number != "006" {
		t.Errorf("await [006] numbers, got [%s]", items[0].Number)
	}
}

func newTestBankAccount(number string) *BankAccount {
	account := NewBankAccount(1, true, number, "USD")
	return &account
}
