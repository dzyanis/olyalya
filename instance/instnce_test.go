package instance

import "testing"

func TestGetSet(t *testing.T) {
	o := New()
	o.Set("foo", "bar")
	if o.Get("foo") != "bar" {
		t.Error("Get")
	}

	o.Set("number", 42)
	if o.Get("number") != 42 {
		t.Error("Get")
	}
}

func TestHas(t *testing.T) {
	o := New()
	if o.Has("somekey") {
		t.Error("Key 'somekey' exists")
	}

	o.Set("somekey", "somekey")
	if !o.Has("somekey") {
		t.Error("Key 'somekey' doesn't exist")
	}
}

func TestLen(t *testing.T) {
	o := New()
	if o.Len() != 0 {
		t.Errorf("Length is %d", o.Len())
	}

	o.Set("test_len", "test_len")
	if o.Len() != 1 {
		t.Errorf("Length is %d", o.Len())
	}

	o.Delete("test_len")
	if o.Len() != 0 {
		t.Errorf("Length is %d", o.Len())
	}
}

func TestKeys(t *testing.T) {
	o := New()
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