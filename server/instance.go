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
	ErrInstanceVariableNotExist = errors.New("Variable is not exist")
	ErrInstanceNotExist = errors.New("Item Not exist")
	//ErrInstanceConvertToInt = errors.New("Convert to int")
	ErrInstanceValueWasExpired = errors.New("Value was expired")

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

func (o *Instance) Set(name string, value interface{}) error {
	o.Lock()
	defer o.Unlock()

	switch reflect.TypeOf(value).Kind() {
	case reflect.Map:
		o.set(name, value)
	case reflect.Slice:
		o.set(name, value)
	case reflect.String:
		o.set(name, value)
	default:
		return ErrInstanceUnknownType
	}

	return nil
}

func (o *Instance) set(k string, v interface{}) error {
	o.base[k] = v
	return nil
}

func (o *Instance) SetTTL(name string, ttl uint) error {
	o.Lock()
	defer o.Unlock()

	_, ok := o.base[name]
	if !ok {
		return ErrInstanceVariableNotExist
	}

	o.setTTL(name, ttl)
	return nil
}

func (o *Instance) Has(name string) bool {
	o.RLock()
	defer o.RUnlock()
	_, ok := o.base[name]
	return ok
}

func (o *Instance) Del(name string) {
	o.Lock()
	defer o.Unlock()

	o.del(name)
}

func (o *Instance) Get(name string) (interface{}, error) {
	o.RLock()
	defer o.RUnlock()
	r, ok := o.base[name]
	if !ok {
		return nil, ErrInstanceVariableNotExist
	}
	if o.checkExpired(name) {
		return nil, ErrInstanceValueWasExpired
	}
	return r, nil
}

func (o *Instance) Len() int {
	o.RLock()
	defer o.RUnlock()
	l := len(o.base)
	return l
}

func (o *Instance) ArrSet(name string, index uint, value string) error {
	o.Lock()
	defer o.Unlock()

	record, ok := o.base[name]
	if !ok {
		return ErrInstanceNotExist
	}
	if o.checkExpired(name) {
		return ErrInstanceValueWasExpired
	}

	s := reflect.ValueOf(record)
	if s.Kind() != reflect.Slice {
		return ErrInstanceValueIsNotArr
	}

	record.([]string)[index] = value

	return nil
}

func (o *Instance) ArrAdd(name string, value string) error {
	o.Lock()
	defer o.Unlock()

	record, ok := o.base[name]
	if !ok {
		return ErrInstanceNotExist
	}
	if o.checkExpired(name) {
		return ErrInstanceValueWasExpired
	}

	s := reflect.ValueOf(record)
	if s.Kind() != reflect.Slice {
		return ErrInstanceValueIsNotArr
	}

	o.base[name] = append(o.base[name].([]string), value)

	return nil
}

func (o *Instance) ArrDel(name string, index uint) error {
	o.Lock()
	defer o.Unlock()

	err := o.arrAccess(name, index)
	if err!=nil {
		return err
	}
	if o.checkExpired(name) {
		return nil
	}

	sl := o.base[name].([]string)
	o.base[name] = append(sl[:index], sl[index+1:]...)

	return nil
}

func (o *Instance) ArrGet(name string, index uint) (string, error) {
	o.RLock()
	defer o.RUnlock()

	err := o.arrAccess(name, index)
	if err!=nil {
		return "", err
	}
	if o.checkExpired(name) {
		return "", ErrInstanceValueWasExpired
	}

	return o.base[name].([]string)[index], nil
}

func (o *Instance) arrAccess(name string, index uint) error {
	record, ok := o.base[name]
	if !ok {
		return ErrInstanceNotExist
	}

	s := reflect.ValueOf(record)
	if s.Kind() != reflect.Slice {
		return ErrInstanceValueIsNotArr
	}

	l := len(record.([]string))
	if index > uint(l)-1 {
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
	if o.checkExpired(name) {
		return ErrInstanceValueWasExpired
	}

	rf := reflect.ValueOf(record)
	if rf.Kind() != reflect.Map {
		return ErrInstanceValueIsNotHash
	}

	record.(map[string]string)[key] = value

	return nil
}

func (o *Instance) del(name string) {
	delete(o.base, name)
	o.delTLL(name)
}

func (o *Instance) checkExpired(name string) bool {
	if o.isExpire(name) {
		o.del(name)
		return true
	}

	return false
}

func (o *Instance) HashDel(name string, key string) error {
	o.Lock()
	defer o.Unlock()

	_, ok := o.base[name]
	if !ok {
		return ErrInstanceNotExist
	}
	if o.checkExpired(name) {
		return nil
	}

	delete(o.base[name].(map[string]string), key)

	return nil
}

func (o *Instance) DelTTL(name string) {
	o.RLock()
	defer o.RUnlock()
	o.delTLL(name)
}

func (o *Instance) HashGet(name string, key string) (string, error) {
	o.RLock()
	defer o.RUnlock()

	record, ok := o.base[name]
	if !ok {
		return "", ErrInstanceNotExist
	}
	if o.checkExpired(name) {
		return "", ErrInstanceValueWasExpired
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
		o.Del(k);
	}
}