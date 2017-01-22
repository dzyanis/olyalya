package main

import (
	"net/http"
	"flag"
	"log"
	"fmt"
	"io/ioutil"
	"reflect"
	"encoding/json"
	"github.com/dzyanis/olyalya/server"
	"github.com/gorilla/pat"
)

var (
	db = server.NewDatabase()
	Version = "0.0.0"
)

var (
	httpAddress = flag.String("http.addr", ":8080", "HTTP listen address")
)

type RespondJson map[string]interface{}

type RequestJsonString struct {
	Key string	`json:"name"`
	Value string	`json:"value"`
	Expire uint	`json:"expire"`
}
type RequestJsonArray struct {
	Key string	`json:"name"`
	Value []string	`json:"value"`
	Expire uint	`json:"expire"`
}
type RequestJsonHash struct {
	Key string		`json:"name"`
	Value map[string]string	`json:"value"`
	Expire uint		`json:"expire"`
}

func main() {
	var router = pat.New();

	router.Post("/db/create", handlerDatabaseCreate)

	router.Post("/db/{instance}/arr/set", handlerInstanceArraySet)
	//router.Post("/db/{instance}/arr/set/{index}", handlerInstanceArraySet)
	//router.Get("/db/{instance}/arr/get/{index}", handlerInstanceArrayGet)
	//router.Delete("/db/{instance}/arr/delete/{key}", handlerInstanceArrayDelete)

	router.Post("/db/{instance}/hash/set", handlerInstanceHashSet)

	router.Post("/db/{instance}/set", handlerInstanceSet)
	router.Get("/db/{instance}/get/{key}", handlerInstanceGet)
	router.Delete("/db/{instance}/delete/{key}", handlerInstanceDelete)

	router.Get("/db/{instance}", handlerInstanceInfo)

	router.Get("/", handlerInfo)
	http.Handle("/", router)

	log.Printf("Version %s listening on %s", Version, *httpAddress)
	log.Fatal(http.ListenAndServe(*httpAddress, nil))
}

func handlerInstanceArraySet(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body);
	defer r.Body.Close()
	if  err != nil {
		handlerJsonError(w, err)
		return
	}

	res := new(RequestJsonArray)
	err = json.Unmarshal([]byte(body), &res)
	if  err != nil {
		handlerJsonError(w, err)
		return
	}

	instanceName := r.URL.Query().Get(":instance")

	if !db.Has(instanceName) {
		handlerJsonError(w, fmt.Errorf("Instance `%s` doesn't exist", instanceName))
	} else {
		db.Get(instanceName).Set(res.Key, res.Value, res.Expire)
		handlerJson(w, http.StatusOK, &RespondJson{
			"status": "OK",
			"exist": db.Has(res.Key),
			"value": res.Value,
			"type": reflect.TypeOf(res.Value),
		})
	}
}

func handlerInstanceHashSet(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body);
	defer r.Body.Close()
	if  err != nil {
		handlerJsonError(w, err)
		return
	}

	res := new(RequestJsonHash)
	err = json.Unmarshal([]byte(body), &res)
	if  err != nil {
		handlerJsonError(w, err)
		return
	}

	instanceName := r.URL.Query().Get(":instance")
	log.Printf("handlerInstanceHashSet: %#v, %s", res, reflect.TypeOf(res.Value))

	if !db.Has(instanceName) {
		err = fmt.Errorf("Instance `%s` doesn't exist", instanceName)
		handlerJsonError(w, err)
	} else {
		db.Get(instanceName).Set(res.Key, res.Value, res.Expire)
		handlerJson(w, http.StatusOK, &RespondJson{
			"status": "OK",
			"exist": db.Has(res.Key),
			"value": res.Value,
			"type": reflect.TypeOf(res.Value),
		})
	}
}

// Create instance of database
// Method: POST
// Content-Type: application/json
// URI: /db/create
// Request: {"name": "string"}
func handlerDatabaseCreate(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{}
	err := json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()

	if err == nil {
		db.Create(data["name"])
		handlerJson(w, http.StatusOK, &RespondJson{
			"status": "OK",
			"body": data,
		})
		return
	}

	handlerJson(w, http.StatusInternalServerError, &RespondJson{
		"status": "ERROR",
		"message": err,
	})
}

func handlerJson(w http.ResponseWriter, code int, object interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(object)
}

func handlerJsonError(w http.ResponseWriter, err error) {
	log.Println(err)
	handlerJson(w, http.StatusInternalServerError, &RespondJson{
		"status": "ERROR",
		"message": err,
	})
}

func handlerJsonOk(w http.ResponseWriter) {
	handlerJson(w, http.StatusOK, &RespondJson{
		"status": "OK",
	})
}

func handlerInfo(w http.ResponseWriter, r *http.Request) {
	handlerJson(w, http.StatusOK, RespondJson{
		"status": "OK",
		"len": db.Len(),
		"keys": db.Keys(),
	})
}

func handlerInstanceInfo(w http.ResponseWriter, r *http.Request) {
	instanceName := r.URL.Query().Get(":instance")
	instance := db.Get(instanceName)
	handlerJson(w, http.StatusOK, &RespondJson{
		"status": "OK",
		"instanceName": instanceName,
		"len": instance.Len(),
		"keys": instance.Keys(),
	})
}

func handlerInstanceGet(w http.ResponseWriter, r *http.Request) {
	instanceName := r.URL.Query().Get(":instance")
	instance := db.Get(instanceName)

	keyName := r.URL.Query().Get(":key")

	value := instance.Get(keyName)
	handlerJson(w, http.StatusOK, &RespondJson{
		"status": "OK",
		"exist": instance.Has(keyName),
		"value": value,
		"type": reflect.TypeOf(value),
	})
}

func handlerInstanceDelete(w http.ResponseWriter, r *http.Request) {
	instanceName := r.URL.Query().Get(":instance")
	instance := db.Get(instanceName)

	keyName := r.URL.Query().Get(":key")
	instance.Delete(keyName)

	handlerJson(w, http.StatusOK, &RespondJson{
		"status": "OK",
	})
}

func handlerInstanceSet(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body);
	defer r.Body.Close()
	if  err != nil {
		handlerJsonError(w, err)
		return
	}

	res := new(RequestJsonString)
	err = json.Unmarshal([]byte(body), &res)
	if  err != nil {
		handlerJsonError(w, err)
		return
	}

	instanceName := r.URL.Query().Get(":instance")
	log.Printf("handlerInstanceSet: %#v, %s", res, reflect.TypeOf(res.Value))

	if !db.Has(instanceName) {
		err = fmt.Errorf("Instance `%s` doesn't exist", instanceName)
		handlerJsonError(w, err)
	} else {
		db.Get(instanceName).Set(res.Key, res.Value, res.Expire)

		handlerJson(w, http.StatusOK, &RespondJson{
			"status": "OK",
			"exist": db.Has(res.Key),
			"value": res.Value,
			"type": reflect.TypeOf(res.Value),
		})
	}
}