package database

import "testing"

func TestCreate(t *testing.T) {
	db := New()
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