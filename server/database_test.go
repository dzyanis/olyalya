package server

import "testing"

func TestCreate(t *testing.T) {
	db := NewDatabase()
	db.Create("dz")
	if db.Len() != 1 {
		t.Error("Database::Create works wrong!")
	}
	db.Create("dz")
	if db.Len() != 1 {
		t.Error("Database::Create works wrong!")
	}
	db.Create("foo")
	db.Create("dz")
	if db.Len() != 2 {
		t.Error("Database::Create works wrong!")
	}
}

func TestGetInstance(t *testing.T) {
	db := NewDatabase()
	db.Create("dz")
	db.Get("dz").Set("one", "1", 0)
	if db.Get("dz").Get("one") != "1" {
		t.Error("Problem with instance")
	}
}