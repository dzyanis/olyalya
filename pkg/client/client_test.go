package client_test

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
	"bytes"
	"io/ioutil"

	client "github.com/dzyanis/olyalya/pkg/client"
)

func TestCreate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/create" {
			t.Error("Unexpected result")
		}

		rb, _ := ioutil.ReadAll(r.Body)
		if !bytes.Equal(rb, []byte(`{"name":"dz"}`)) {
			t.Error("Unexpected result")
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"OK"}`)
	}))
	defer ts.Close()

	cln := client.NewClient(ts.URL)
	err := cln.CreateInstance("dz")
	if err != nil {
		t.Error(err)
	}
}

func TestSelectUnexist(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/in/unexist" {
			t.Error("Unexpected result")
		}

		rb, _ := ioutil.ReadAll(r.Body)
		if !bytes.Equal(rb, []byte(`{}`)) {
			t.Error("Unexpected result")
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"error":"Instance not exist","status":"ERROR"}`)
	}))
	defer ts.Close()

	cln := client.NewClient(ts.URL)
	err := cln.SelectInstance("unexist")
	if err == nil {
		t.Error("Unexpected result")
	}
}

func TestSelectOk(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/in/dz" {
			t.Error("Unexpected result")
		}

		rb, _ := ioutil.ReadAll(r.Body)
		if !bytes.Equal(rb, []byte(`{}`)) {
			t.Error("Unexpected result")
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"instance":"dz","length":0,"names":[]}`)
	}))
	defer ts.Close()

	cln := client.NewClient(ts.URL)
	err := cln.SelectInstance("dz")
	if err != nil {
		t.Error(err)
	}
}

//func TestGetUnexist(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		fmt.Println(r.RequestURI)
//		if r.RequestURI != "/in/dz" {
//			t.Error("Unexpected result")
//		}
//
//		rb, _ := ioutil.ReadAll(r.Body)
//		if !bytes.Equal(rb, []byte(`{}`)) {
//			t.Error("Unexpected result")
//		}
//
//		w.WriteHeader(http.StatusOK)
//		fmt.Fprintf(w, `{"error":"Variable is not exist","status":"ERROR"}`)
//	}))
//	defer ts.Close()
//
//	cln := NewClient(ts.URL)
//	cln.SelectInstance("dz")
//	_, err := cln.Get("one")
//	if err == nil {
//		t.Error("Unexpected result")
//	}

	//err = cln.Set("one", "1", 0)
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//one, err = cln.Get("one")
	//if err != nil {
	//	t.Error(err)
	//}
	//if one != "1" {
	//	t.Error("Unexpected result")
	//}
//}
//
//func TestGetSetArray(t *testing.T) {
//	numbers, err := Cln.GetArray("numbers")
//	if err == nil {
//		t.Error("Unexpected result")
//	}
//
//	primary := []string{"zero", "one", "two"}
//	err = Cln.SetArray("numbers", primary, 0)
//	if err != nil {
//		t.Error(err)
//	}
//
//	numbers, err = Cln.GetArray("numbers")
//	if err != nil {
//		t.Error(err)
//	}
//	if len(numbers) != len(primary) {
//		t.Error("Unexpected result")
//	}
//}
//
//func TestArrayElementsGet(t *testing.T) {
//	two, err := Cln.GetArrayElement("numbers", 2)
//	if err != nil {
//		t.Error(err)
//	}
//	if two != "two" {
//		t.Error("Unexpected result")
//	}
//
//	_, err = Cln.GetArrayElement("numbers", 3)
//	if err == nil {
//		t.Error("Unexpected result")
//	}
//}
//
//func TestArrayElementsAdd(t *testing.T) {
//	err := Cln.AddArrayElement("numbers", "three")
//	if err != nil {
//		t.Error(err)
//	}
//	three, err := Cln.GetArrayElement("numbers", 3)
//	if err != nil {
//		t.Error(err)
//	}
//	if three != "three" {
//		t.Error("Unexpected result")
//	}
//}
//
//func TestArrayElementsSet(t *testing.T) {
//	err := Cln.SetArrayElement("numbers", 0, "1")
//	if err != nil {
//		t.Error(err)
//	}
//	three, err := Cln.GetArrayElement("numbers", 3)
//	if err != nil {
//		t.Error(err)
//	}
//	if three != "three" {
//		t.Error("Unexpected result")
//	}
//}
//
//func TestArrayElementsDel(t *testing.T) {
//	err := Cln.DelArrayElement("numbers", 0)
//	if err != nil {
//		t.Error(err)
//	}
//
//	one, err := Cln.GetArrayElement("numbers", 0)
//	if err != nil {
//		t.Error(err)
//	}
//	if one != "one" {
//		t.Error(one)
//		t.Error("Unexpected result")
//	}
//}
//
//func TestGetSetHash(t *testing.T) {
//	user, err := Cln.GetHash("user")
//	if err == nil {
//		t.Error("Unexpected result")
//	}
//
//	primary := map[string]string{
//		"name":     "Dzyanis Kuzmenka",
//		"language": "C",
//	}
//	err = Cln.SetHash("user", primary, 0)
//	if err != nil {
//		t.Error(err)
//	}
//
//	user, err = Cln.GetHash("user")
//	if err != nil {
//		t.Error(err)
//	}
//	if !reflect.DeepEqual(user, primary) {
//		t.Error("Unexpected result")
//	}
//}
//
//func TestHashSetGetElement(t *testing.T) {
//	lang, err := Cln.GetHashElement("user", "language")
//	if err != nil {
//		t.Error(err)
//	}
//	if lang != "C" {
//		t.Error("Unexpected result")
//	}
//
//	err = Cln.SetHashElement("user", "language", "Go")
//	if err != nil {
//		t.Error(err)
//	}
//
//	lang, err = Cln.GetHashElement("user", "language")
//	if err != nil {
//		t.Error(err)
//	}
//	if lang != "Go" {
//		t.Error("Unexpected result")
//	}
//}
//
//func TestHashAddDelElement(t *testing.T) {
//	err := Cln.SetHashElement("user", "year", "1987")
//	if err != nil {
//		t.Error(err)
//	}
//
//	year, err := Cln.GetHashElement("user", "year")
//	if err != nil {
//		t.Error(err)
//	}
//	if year != "1987" {
//		t.Error("Unexpected result")
//	}
//
//	err = Cln.DelHashElement("user", "year")
//	if err != nil {
//		t.Error(err)
//	}
//
//	_, err = Cln.GetHashElement("user", "year")
//	if err == nil {
//		t.Error("Unexpected result")
//	}
//}

func TestDestroy(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		if r.RequestURI != "/destroy" {
			t.Error("Unexpected result")
		}

		rb, _ := ioutil.ReadAll(r.Body)
		if !bytes.Equal(rb, []byte(`{"name":"dz"}`)) {
			t.Error(string(rb))
			t.Error("Unexpected result")
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"OK"}`)
	}))
	defer ts.Close()

	cln := client.NewClient(ts.URL)
	err := cln.Destroy("dz")
	if err != nil {
		t.Error(err)
	}
}
