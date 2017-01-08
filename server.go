package main

import (
	"net/http"
	"flag"
	"log"
	"encoding/json"
	"github.com/dzyanis/olyalya/database"
)

var (
	db = database.New()
	Version = "0.0"
)

var (
	httpAddress = flag.String("http.addr", ":8080", "HTTP listen address")
)

func main() {
	http.HandleFunc("/", handlerInfo)
	http.HandleFunc("/db/create", handlerDatabaseCreate)

	//http.HandleFunc("/set", handlerSet)
	//http.HandleFunc("/get", handlerGet)
	//http.HandleFunc("/delete", handlerDelete)

	log.Printf("Version %s listening on %s", Version, *httpAddress)
	log.Fatal(http.ListenAndServe(*httpAddress, nil))
}

func handlerDatabaseCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	data := map[string]string{}
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

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

//func handlerSet(w http.ResponseWriter, r *http.Request) {
//	//key   := r.URL.Query().Get("key")
//	//value := r.URL.Query().Get("value")
//	//Base.Set(key, value)
//	request := map[string]interface{}{
//		"status": "OK",
//	}
//	handlerJson(w, http.StatusOK, request)
//}
//
//func handlerGet(w http.ResponseWriter, r *http.Request) {
//	key := r.URL.Query().Get("key")
//	fmt.Fprintf(w, "value: %+v!", Base.Get(key))
//}
//
//func handlerDelete(w http.ResponseWriter, r *http.Request) {
//	flag.Parse()
//	key   := r.URL.Query().Get("key")
//	Base.Delete(key)
//	fmt.Fprint(w, "OK!")
//}