package server

import (
	"sync"
	"errors"
	"reflect"
)

type (
	InstanceTypeBase map[string]interface{}
)

type Instance struct {
	sync.RWMutex
	base InstanceTypeBase
	ExtensionExpire
}

var (
	ErrInstanceUnknownType = errors.New("Unknown type")
	ErrInstanceNotExist = errors.New("Item Not exist")
	ErrInstanceConvertToInt = errors.New("Convert to int")

	ErrInstanceValueIsNotArr = errors.New("Value is not array")
	ErrInstanceArrIndexIsNotExist = errors.New("Index is not exist")

	ErrInstanceValueIsNotHash = errors.New("Value is not hash")
	ErrInstanceHashKeyIsNotExist = errors.New("Key is not exist")
)

func NewInstance() *Instance {
	return &Instance{
		base: make(InstanceTypeBase),
		ExtensionExpire: NewExtensionExpire(),
	}
}

func (o *Instance) Set(key string, value interface{}, params... interface{}) error {
	o.Lock()
	defer o.Unlock()

	switch reflect.TypeOf(value).Kind() {
	case reflect.Map:
		o.set(key, value)
	case reflect.Slice:
		o.set(key, value)
	case reflect.String:
		o.set(key, value)
	default:
		return ErrInstanceUnknownType
	}

	l := len(params)
	if l > 0 {
		// @todo looks ugly
		u, ok := params[0].(int)
		if !ok {
			return ErrInstanceConvertToInt
		}
		o.setTTL(key, uint(u))
	}

	return nil
}

func (o *Instance) set(k string, v interface{}) error {
	o.base[k] = v
	return nil
}

func (o *Instance) Has(key string) bool {
	o.RLock()
	defer o.RUnlock()
	_, ok := o.base[key]
	return ok
}

func (o *Instance) Delete(key string) {
	o.Lock()
	defer o.Unlock()
	delete(o.base, key)
	o.DeleteTLL(key)
}

func (o *Instance) Get(key string) (interface{}) {
	o.RLock()
	defer o.RUnlock()
	r, ok := o.base[key]
	if !ok {
		return nil
	}
	return r
}

func (o *Instance) Len() int {
	o.RLock()
	defer o.RUnlock()
	l := len(o.base)
	return l
}

func (o *Instance) ArrSet(key string, index uint, value string) error {
	o.Lock()
	defer o.Unlock()

	record, ok := o.base[key]
	if !ok {
		return ErrInstanceNotExist
	}

	s := reflect.ValueOf(record)
	if s.Kind() != reflect.Slice {
		return ErrInstanceValueIsNotArr
	}

	record.([]string)[index] = value

	return nil
}

func (o *Instance) ArrAdd(key string, value string) error {
	o.Lock()
	defer o.Unlock()

	record, ok := o.base[key]
	if !ok {
		return ErrInstanceNotExist
	}

	s := reflect.ValueOf(record)
	if s.Kind() != reflect.Slice {
		return ErrInstanceValueIsNotArr
	}

	o.base[key] = append(o.base[key].([]string), value)

	return nil
}

func (o *Instance) ArrDel(key string, index uint) error {
	o.Lock()
	defer o.Unlock()

	err := o.arrAccess(key, index)
	if err!=nil {
		return err
	}

	sl := o.base[key].([]string)
	o.base[key] = append(sl[:index], sl[index+1:]...)

	return nil
}

func (o *Instance) ArrGet(key string, index uint) (string, error) {
	o.RLock()
	defer o.RUnlock()

	err := o.arrAccess(key, index)
	if err!=nil {
		return "", err
	}

	return o.base[key].([]string)[index], nil
}

func (o *Instance) arrAccess(key string, index uint) error {
	record, ok := o.base[key]
	if !ok {
		return ErrInstanceNotExist
	}

	s := reflect.ValueOf(record)
	if s.Kind() != reflect.Slice {
		return ErrInstanceValueIsNotArr
	}

	l := len(record.([]string))
	if uint(l) < index {
		return ErrInstanceArrIndexIsNotExist
	}

	return nil
}

func (o *Instance) HashSet(name string, key string, value string) error {
	o.Lock()
	defer o.Unlock()

	record, ok := o.base[name]
	if !ok {
		return ErrInstanceNotExist
	}

	rf := reflect.ValueOf(record)
	if rf.Kind() != reflect.Map {
		return ErrInstanceValueIsNotHash
	}

	record.(map[string]string)[key] = value

	return nil
}

func (o *Instance) HashDel(name string, key string) error {
	o.Lock()
	defer o.Unlock()

	_, ok := o.base[name]
	if !ok {
		return ErrInstanceNotExist
	}

	delete(o.base[name].(map[string]string), key)

	return nil
}

func (o *Instance) HashGet(name string, key string) (string, error) {
	o.RLock()
	defer o.RUnlock()

	record, ok := o.base[name]
	if !ok {
		return "", ErrInstanceNotExist
	}

	rf := reflect.ValueOf(record)
	if rf.Kind() != reflect.Map {
		return "", ErrInstanceValueIsNotHash
	}

	val, ok := record.(map[string]string)[key]
	if !ok {
		return "", ErrInstanceHashKeyIsNotExist
	}

	return val, nil
}

// It work so slow
func (o *Instance) Keys() []string {
	o.RLock()
	defer o.RUnlock()
	keys := make([]string, 0, o.Len())
	for k := range o.base {
		keys = append(keys, k)
	}
	return keys
}

func (o *Instance) GetExpiredKeys() []string {
	o.RLock()
	defer o.RUnlock()
	return o.getExpiredKeys()
}

func (o *Instance) Cleaner() {
	for _, k := range o.GetExpiredKeys() {
		o.Delete(k);
	}
}