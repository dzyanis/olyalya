package server

import (
	"errors"
)

type DataBaseTypeInstances map[string]*Instance

type DataBase struct {
	instances DataBaseTypeInstances
}

var (
	ErrDatabaseInstanceNotExist = errors.New("Instance Not Exist")
)

func NewDatabase() *DataBase {
	return &DataBase{
		instances: make(DataBaseTypeInstances),
	}
}

func (db *DataBase) Create(name string) {
	if !db.Has(name) {
		db.instances[name] = NewInstance()
	}
}

func (db DataBase) Has(key string) bool {
	_, ok := db.instances[key]
	return ok
}

func (db DataBase) Delete(key string) {
	delete(db.instances, key)
}

func (db DataBase) Get(key string) (*Instance, error) {
	v, ok := db.instances[key]
	if !ok {
		return nil, ErrDatabaseInstanceNotExist
	}
	return v, nil
}

func (db DataBase) Len() int {
	return len(db.instances)
}

func (db DataBase) Keys() []string {
	keys := make([]string, 0, db.Len())
	for k := range db.instances {
		keys = append(keys, k)
	}
	return keys
}