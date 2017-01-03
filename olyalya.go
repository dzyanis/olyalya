package olyalya

import "sync"

type OlyalyaTypeBase map[string]interface{}

type Olyalya struct {
	sync.RWMutex
	base OlyalyaTypeBase
}

func InitOlyalya() Olyalya {
	return Olyalya{base: make(OlyalyaTypeBase)}
}

func (o Olyalya) Set(key string, value interface{}) {
	o.Lock()
	defer o.Unlock()
	o.base[key] = value
}

func (o Olyalya) Has(key string) bool {
	o.RLock()
	defer o.RUnlock()
	_, ok := o.base[key]
	return ok
}

func (o Olyalya) Delete(key string) {
	o.Lock()
	defer o.Unlock()
	delete(o.base, key)
}

func (o Olyalya) Get(key string) interface{} {
	o.RLock()
	defer o.RUnlock()
	v, _ := o.base[key]
	return v
}

func (o Olyalya) Len() int {
	o.RLock()
	defer o.RUnlock()
	l := len(o.base)
	return l
}

func (o Olyalya) Keys() []string {
	o.RLock()
	defer o.RUnlock()
	keys := make([]string, 0, o.Len())
	for k := range o.base {
		keys = append(keys, k)
	}
	return keys
}