package subdomain

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

	aTitle := "a"

	aId, err := service.Create(Subdomain{
		Title: aTitle,
	})
	if err != nil {
		t.Error(err)
	}

	aModel, err := service.Describe(aId)
	if err != nil {
		t.Error(err)
	}
	if aModel.Title != aTitle {
		t.Errorf("wrong model title returns %s != %s", aModel.Title, aTitle)
	}

	aNewTitle := "a - new"

	err = service.Update(aId, Subdomain{Title: aNewTitle})
	if err != nil {
		t.Error(err)
	}

	aModel, err = service.Describe(aId)
	if err != nil {
		t.Error(err)
	}
	if aModel.Title != aNewTitle {
		t.Errorf("wrong model title returns %s != %s", aModel.Title, aNewTitle)
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

	aModel, err = service.Describe(aId)
	if err == nil {
		t.Error("await error, nil returns")
	}
	if nil != aModel {
		t.Error("await nil, model returns")
	}
}

func TestServiceList(t *testing.T) {
	service := NewService()

	titles := []string{
		"first",
		"second",
		"thirds",
		"fourths",
		"fifth",
	}

	titleIds := make(map[string]uint64)

	for _, title := range titles {
		id, err := service.Create(Subdomain{Title: title})
		if err != nil {
			t.Errorf("error when create new: %v", err)
		}
		titleIds[title] = id
	}

	items, err := service.List(0, 2)
	if err != nil {
		t.Errorf("error when get list: %v", err)
	}
	if len(items) != 2 {
		t.Errorf("await list len 2, got %d items", len(items))
	}
	if items[0].Title != "first" && items[1].Title != "second" {
		t.Errorf("await [first,second] titles, got [%s,%s]", items[0].Title, items[1].Title)
	}

	items, err = service.List(items[1].GetId(), 3)

	if err != nil {
		t.Errorf("error when get list: %v", err)
	}
	if len(items) != 3 {
		t.Errorf("await list len 3, got %d items", len(items))
	}
	if items[0].Title != "thirds" && items[1].Title != "fourths" && items[2].Title != "fifth" {
		t.Errorf("await [thirds,fourths,fifth] titles, got [%s,%s,%s]", items[0].Title, items[1].Title, items[2].Title)
	}

	items, err = service.List(items[2].GetId(), 10)
	if err == nil {
		t.Error("await error, got nothing")
	}

	toRemoveTitles := []string{"thirds", "fifth", "first"}
	for _, title := range toRemoveTitles {
		id := titleIds[title]

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
	if items[0].Title != "second" && items[1].Title != "fourths" {
		t.Errorf("await [second,fourths] titles, got [%s,%s]", items[0].Title, items[1].Title)
	}
}
