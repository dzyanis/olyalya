package client

import (
	"flag"
	"reflect"
	"testing"
)

var (
	httpUrl  = flag.String("http.url", "localhost", "HTTP listen URL")
	httpPort = flag.Int("http.port", 3000, "HTTP listen port")

	Cln *Client
)

func init() {
	flag.Parse()
	Cln = NewClient(*httpUrl, *httpPort)
}

func TestCreate(t *testing.T) {
	err := Cln.CreateInstance("dz")
	if err != nil {
		t.Error(err)
	}

	err = Cln.CreateInstance("dz")
	if err == nil {
		t.Error("Unexpected result")
	}
}

func TestSelect(t *testing.T) {
	err := Cln.SelectInstance("unexist")
	if err == nil {
		t.Error("Unexpected result")
	}

	err = Cln.SelectInstance("dz")
	if err != nil {
		t.Error(err)
	}
}

func TestGetSet(t *testing.T) {
	one, err := Cln.Get("one")
	if err == nil {
		t.Error("Unexpected result")
	}

	err = Cln.Set("one", "1", 0)
	if err != nil {
		t.Error(err)
	}

	one, err = Cln.Get("one")
	if err != nil {
		t.Error(err)
	}
	if one != "1" {
		t.Error("Unexpected result")
	}
}

func TestGetSetArray(t *testing.T) {
	numbers, err := Cln.GetArray("numbers")
	if err == nil {
		t.Error("Unexpected result")
	}

	primary := []string{"zero", "one", "two"}
	err = Cln.SetArray("numbers", primary, 0)
	if err != nil {
		t.Error(err)
	}

	numbers, err = Cln.GetArray("numbers")
	if err != nil {
		t.Error(err)
	}
	if len(numbers) != len(primary) {
		t.Error("Unexpected result")
	}
}

func TestArrayElementsGet(t *testing.T) {
	two, err := Cln.GetArrayElement("numbers", 2)
	if err != nil {
		t.Error(err)
	}
	if two != "two" {
		t.Error("Unexpected result")
	}

	_, err = Cln.GetArrayElement("numbers", 3)
	if err == nil {
		t.Error("Unexpected result")
	}
}

func TestArrayElementsAdd(t *testing.T) {
	err := Cln.AddArrayElement("numbers", "three")
	if err != nil {
		t.Error(err)
	}
	three, err := Cln.GetArrayElement("numbers", 3)
	if err != nil {
		t.Error(err)
	}
	if three != "three" {
		t.Error("Unexpected result")
	}
}

func TestArrayElementsSet(t *testing.T) {
	err := Cln.SetArrayElement("numbers", 0, "1")
	if err != nil {
		t.Error(err)
	}
	three, err := Cln.GetArrayElement("numbers", 3)
	if err != nil {
		t.Error(err)
	}
	if three != "three" {
		t.Error("Unexpected result")
	}
}

func TestArrayElementsDel(t *testing.T) {
	err := Cln.DelArrayElement("numbers", 0)
	if err != nil {
		t.Error(err)
	}

	one, err := Cln.GetArrayElement("numbers", 0)
	if err != nil {
		t.Error(err)
	}
	if one != "one" {
		t.Error(one)
		t.Error("Unexpected result")
	}
}

func TestGetSetHash(t *testing.T) {
	user, err := Cln.GetHash("user")
	if err == nil {
		t.Error("Unexpected result")
	}

	primary := map[string]string{
		"name":     "Dzyanis Kuzmenka",
		"language": "C",
	}
	err = Cln.SetHash("user", primary, 0)
	if err != nil {
		t.Error(err)
	}

	user, err = Cln.GetHash("user")
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(user, primary) {
		t.Error("Unexpected result")
	}
}

func TestHashSetGetElement(t *testing.T) {
	lang, err := Cln.GetHashElement("user", "language")
	if err != nil {
		t.Error(err)
	}
	if lang != "C" {
		t.Error("Unexpected result")
	}

	err = Cln.SetHashElement("user", "language", "Go")
	if err != nil {
		t.Error(err)
	}

	lang, err = Cln.GetHashElement("user", "language")
	if err != nil {
		t.Error(err)
	}
	if lang != "Go" {
		t.Error("Unexpected result")
	}
}

func TestHashAddDelElement(t *testing.T) {
	err := Cln.SetHashElement("user", "year", "1987")
	if err != nil {
		t.Error(err)
	}

	year, err := Cln.GetHashElement("user", "year")
	if err != nil {
		t.Error(err)
	}
	if year != "1987" {
		t.Error("Unexpected result")
	}

	err = Cln.DelHashElement("user", "year")
	if err != nil {
		t.Error(err)
	}

	_, err = Cln.GetHashElement("user", "year")
	if err == nil {
		t.Error("Unexpected result")
	}
}

func TestDestroy(t *testing.T) {
	err := Cln.Destroy("dz")
	if err != nil {
		t.Error(err)
	}
}
