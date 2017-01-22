package server

import (
	"testing"
	"reflect"
)

func TestGetSet(t *testing.T) {
	o := NewInstance()
	o.Set("foo", "bar", 0)
	if o.Get("foo") != "bar" {
		t.Error("Method Get returned unexpected result:", o.Get("foo"))
	}

	o.Set("number", "42", 0)
	if o.Get("number") != "42" {
		t.Error("Method Get returned unexpected result:", o.Get("number"))
	}
}

func TestGetSetArr(t *testing.T) {
	o := NewInstance()

	arr := []string{"one", "two", "three"}
	o.Set("arr", arr, 0)
	as := o.Get("arr");
	if reflect.TypeOf(as).String() != "[]string" {
		t.Errorf("Function has returned the wrong type %s", reflect.TypeOf(as))
	}
}

func TestHas(t *testing.T) {
	o := NewInstance()
	if o.Has("somekey") {
		t.Error("Key 'somekey' exists")
	}

	o.Set("somekey", "somekey", 0)
	if !o.Has("somekey") {
		t.Error("Key 'somekey' doesn't exist")
	}
}

func TestLen(t *testing.T) {
	o := NewInstance()
	if o.Len() != 0 {
		t.Errorf("Length is %d", o.Len())
	}

	o.Set("test_len", "test_len", 0)
	if o.Len() != 1 {
		t.Errorf("Length is %d", o.Len())
	}

	o.Delete("test_len")
	if o.Len() != 0 {
		t.Errorf("Length is %d", o.Len())
	}
}

func TestKeys(t *testing.T) {
	o := NewInstance()
	ts := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
	}

	for k, v := range ts {
		o.Set(k, v, 0)
	}

	for _, j := range o.Keys() {
		if !o.Has(j) {
			t.Errorf("Key '%s' was found", j)
		}
	}
}