package database

import (
	"github.com/dzyanis/olyalya/instance"
)

type DataBase struct {
	instances map[string]instance.Instance
}

func New() *DataBase {
	return &DataBase{
		instances: make(map[string]instance.Instance),
	}
}

func (db *DataBase) Create(name string) {
	if !db.Has(name) {
		db.instances[name] = instance.New()
	}
}

func (db DataBase) Has(key string) bool {
	_, ok := db.instances[key]
	return ok
}

func (db DataBase) Delete(key string) {
	delete(db.instances, key)
}

func (db DataBase) Get(key string) interface{} {
	v, _ := db.instances[key]
	return v
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