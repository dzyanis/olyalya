package server

import (
	"testing"
	"reflect"
)

func TestGetSet(t *testing.T) {
	o := NewInstance()

	// string
	bar := "bar"
	o.Set("foo", bar)
	v, _ := o.Get("foo")
	if v != bar {
		t.Error("Method Get returned unexpected result: %#v != %#v", bar, v, reflect.TypeOf(v))
	}

	// string
	o.Set("number", "42")
	v, _ = o.Get("number")
	if v != "42" {
		t.Error("Method Get returned unexpected result:", v)
	}

	// Number
	err := o.Set("number", 42)
	if err==nil {
		t.Error("Method Set expected error")
	}
}

func TestGetSetArr(t *testing.T) {
	o := NewInstance()

	arr := []string{"one", "two", "three"}
	o.Set("arr", arr)
	as, _ := o.Get("arr")
	if reflect.TypeOf(as).String() != "[]string" {
		t.Errorf("Function has returned the wrong type %s", reflect.TypeOf(as))
	}
}

func TestHas(t *testing.T) {
	o := NewInstance()
	if o.Has("somekey") {
		t.Error("Key 'somekey' exists")
	}

	o.Set("somekey", "somekey")
	if !o.Has("somekey") {
		t.Error("Key 'somekey' doesn't exist")
	}
}

func TestLen(t *testing.T) {
	o := NewInstance()
	if o.Len() != 0 {
		t.Errorf("Length is %d", o.Len())
	}

	o.Set("test_len", "test_len")
	if o.Len() != 1 {
		t.Errorf("Length is %d", o.Len())
	}

	o.Del("test_len")
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
		o.Set(k, v)
	}

	for _, j := range o.Keys() {
		if !o.Has(j) {
			t.Errorf("Key '%s' was found", j)
		}
	}
}

func TestArrayGet(t *testing.T) {
	o := NewInstance()
	o.Set("arr", []string{"0", "1", "2", "3"})

	_, err := o.ArrGet("arr", 4)
	if err==nil {
		t.Error("Error was expected")
	}

	two, err := o.ArrGet("arr", 3)
	if two!="3" {
		t.Error("Wrong value")
	}

	o.Set("string", "ELLO GOVNA")
	_, err = o.ArrGet("string", 0)
	if err==nil {
		t.Error("Error was expected")
	}
}

func TestArrayDel(t *testing.T) {
	o := NewInstance()
	o.Set("arr", []string{"0", "1", "2", "3"})

	o.ArrDel("arr", 2)
	three, err := o.ArrGet("arr", 2)
	if err!=nil {
		t.Error(err)
	}
	if three!="3" {
		t.Error("Wrong value")
	}
}

func TestArraySetAdd(t *testing.T) {
	o := NewInstance()
	o.Set("arr", []string{"0", "1", "2", "3"})

	o.ArrAdd("arr", "four")
	arr2, _ := o.Get("arr")
	if len(arr2.([]string))!=5 {
		t.Error("Length should be 4")
	}

	o.ArrSet("arr", 4, "four")
	arr2, _ = o.Get("arr")
	two, err := o.ArrGet("arr", 2)
	if err!=nil {
		t.Error(err)
	}
	if two!=arr2.([]string)[2] {
		t.Errorf("Values is not equiel: '%v' != '%v'", two, arr2.([]string)[2])
	}
}

func TestHashGet(t *testing.T) {
	o := NewInstance()
	o.Set("hash", map[string]string{
		"zero": "0",
		"one": "1",
		"two": "2",
		"three": "3",
	})

	_, err := o.HashGet("notexistname", "notexistkey")
	if err==nil {
		t.Error("Error was expected")
	}

	_, err = o.HashGet("hash", "notexistkey")
	if err==nil {
		t.Error("Error was expected")
	}

	one, err := o.HashGet("hash", "one")
	if one!="1" {
		t.Error("Wrong value")
	}

	o.Set("string", "ELLO GOVNA")
	_, err = o.HashGet("string", "notexistkey")
	if err==nil {
		t.Error("Error was expected")
	}
}

func TestHashSet(t *testing.T) {
	o := NewInstance()
	o.Set("hash", map[string]string{
		"zero": "0",
		"one": "1",
		"two": "2",
		"three": "3",
	})

	err := o.HashSet("notexistname", "three", "4")
	if err==nil {
		t.Error("Error was expected")
	}

	err = o.HashSet("hash", "four", "4")
	four, _ := o.HashGet("hash", "four")
	if four!="4" {
		t.Error("Wrong value")
	}

	_ = o.HashSet("hash", "four", "IV")
	four, _ = o.HashGet("hash", "four")
	if four!="IV" {
		t.Error("Wrong value")
	}
}

func TestHashDel(t *testing.T) {
	o := NewInstance()
	o.Set("hash", map[string]string{
		"zero": "0",
		"one": "1",
		"two": "2",
		"three": "3",
	})

	err := o.HashDel("notexistname", "notexistkey")
	if err==nil {
		t.Error("Error was expected")
	}

	err = o.HashDel("hash", "notexistkey")
	if err!=nil {
		t.Error("Error was not expected")
	}

	err = o.HashDel("hash", "two")
	if err!=nil {
		t.Error("Error was not expected")
	}
	_, err = o.HashGet("hash", "two")
	if err==nil {
		t.Error("Error was expected")
	}
}