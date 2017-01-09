package main

import (
	"net/http"
	"flag"
	"log"
	"encoding/json"
	"github.com/dzyanis/olyalya/database"
	"github.com/gorilla/pat"
)

var (
	db = database.New()
	Version = "0.0"
)

var (
	httpAddress = flag.String("http.addr", ":8080", "HTTP listen address")
)

func main() {
	var router = pat.New();

	router.Post("/db/create", handlerDatabaseCreate)

	router.Post("/db/{instance}/set", handlerInstanceSet)
	router.Get("/db/{instance}/get/{key}", handlerInstanceGet)
	router.Get("/db/{instance}", handlerInstanceInfo)

	router.Get("/", handlerInfo)
	http.Handle("/", router)

	log.Printf("Version %s listening on %s", Version, *httpAddress)
	log.Fatal(http.ListenAndServe(*httpAddress, nil))
}

func handlerDatabaseCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	data := map[string]string{}
	if err := decoder.Decode(&data); err != nil {
		// error
	}
	db.Create(data["name"])

	request := map[string]interface{}{
		"status": "OK",
		"body": data,
	}
	handlerJson(w, 200, request)
}

func handlerJson(w http.ResponseWriter, code int, object interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(object)
}

func handlerInfo(w http.ResponseWriter, r *http.Request) {
	request := map[string]interface{}{
		"status": "OK",
		"len": db.Len(),
		"keys": db.Keys(),
	}
	handlerJson(w, http.StatusOK, request)
}

func handlerInstanceInfo(w http.ResponseWriter, r *http.Request) {
	instanceName := r.URL.Query().Get(":instance")
	instance := db.Get(instanceName)
	request := map[string]interface{}{
		"status": "OK",
		"instanceName": instanceName,
		"len": instance.Len(),
		"keys": instance.Keys(),
	}
	handlerJson(w, http.StatusOK, request)
}

func handlerInstanceGet(w http.ResponseWriter, r *http.Request) {
	instanceName := r.URL.Query().Get(":instance")
	instance := db.Get(instanceName)

	keyName := r.URL.Query().Get(":key")

	request := map[string]interface{}{
		"status": "OK",
		"exist": instance.Has(keyName),
		"value": instance.Get(keyName),
	}
	handlerJson(w, http.StatusOK, request)
}

func handlerInstanceSet(w http.ResponseWriter, r *http.Request) {
	instanceName := r.URL.Query().Get(":instance")
	instance := db.Get(instanceName)

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	data := map[string]string{}
	if err := decoder.Decode(&data); err != nil {
		// error
	}

	for k, v := range data {
		instance.Set(k, v)
	}

	request := map[string]interface{}{
		"status": "OK",
		"data": data,
	}
	handlerJson(w, http.StatusOK, request)
}