package database

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
	instance, err := db.Get("dz")
	if err!=nil {
		t.Error(err)
	}
	instance.Set("one", "1")

	instance, err = db.Get("dz")
	if err!=nil {
		t.Error(err)
	}

	v, _ := instance.Get("one")
	if  v!= "1" {
		t.Error("Problem with instance")
	}
}