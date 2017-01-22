package server

import (
	"sync"
)

type InstanceTypeBase map[string]interface{}

type Instance struct {
	sync.RWMutex
	base InstanceTypeBase
	ExtensionExpire
}

func NewInstance() *Instance {
	return &Instance{
		base: make(InstanceTypeBase),
		ExtensionExpire: NewExtensionExpire(),
	}
}

func (o *Instance) Set(key string, value interface{}, ttl uint) {
	o.Lock()
	defer o.Unlock()

	switch value.(type) {
	case map[string]string:
		o.base[key] = value
		o.SetTTL(key, ttl)
	case []string:
		o.base[key] = value
		o.SetTTL(key, ttl)
	case string:
		o.base[key] = value
		o.SetTTL(key, ttl)
	}
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

func (o *Instance) Get(key string) interface{} {
	o.RLock()
	defer o.RUnlock()
	v, _ := o.base[key]
	return v
}

func (o *Instance) Len() int {
	o.RLock()
	defer o.RUnlock()
	l := len(o.base)
	return l
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

func (o *Instance) Cleaner() {
	o.RLock()
	defer o.RUnlock()

	for k := range o.ttl {
		o.CheckExpired(k)
	}
}