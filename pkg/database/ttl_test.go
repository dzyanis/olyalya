package database

import (
	"testing"
	"time"
)

func TestSetGetTTL(t *testing.T) {
	o := NewExtensionExpire()

	o.setTTL("ten", 10)
	if o.GetTTL("ten") > CurrentUnixTime()+10 {
		t.Error("Method Get returned unexpected result:", o.GetTTL("ten"))
	}
}

func TestHasTTL(t *testing.T) {
	o := NewExtensionExpire()

	if o.hasTTL("somekey") {
		t.Error("Key 'somekey' exists")
	}

	o.setTTL("somekey", 10)
	if !o.hasTTL("somekey") {
		t.Error("Key 'somekey' doesn't exist")
	}
}

func TestIsExpireTTL(t *testing.T) {
	o := NewExtensionExpire()

	o.setTTL("2sec", 2)
	if o.isExpire("2sec") {
		t.Errorf("Now '%d'. Date '%d' was expired.", CurrentUnixTime(), o.GetTTL("2sec"))
	}
	time.Sleep(2*time.Second)
	if !o.isExpire("2sec") {
		t.Errorf("Now '%d'. Date '%d' doesn't expired.", CurrentUnixTime(), o.GetTTL("2sec"))
	}
}
