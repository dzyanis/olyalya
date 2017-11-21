// This O(lya-lya) DataBase
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dzyanis/olyalya/pkg/database"
	"github.com/gorilla/pat"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	Db         = database.NewDatabase()
	buildName  = ""
	ServerName = "OLL"
)

var (
	httpAddress = flag.String("http.addr", ":3000", "HTTP listen address")
	version     = flag.Bool("version", false, "Print a version")
)

type RespondJson map[string]interface{}

type RequestJsonTTL struct {
	Name   string `json:"name"`
	Expire uint   `json:"ttl"`
}

type RequestJsonString struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Expire uint   `json:"ttl"`
}

type RequestJsonArray struct {
	Name   string   `json:"name"`
	Value  []string `json:"value"`
	Index  int      `json:"index"`
	Expire uint     `json:"ttl"`
}

type RequestJsonArrayItem struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Index uint   `json:"index"`
}

type RequestJsonHash struct {
	Name   string            `json:"name"`
	Value  map[string]string `json:"value"`
	Key    string            `json:"key"`
	Expire uint              `json:"ttl"`
}

type RequestJsonHashItem struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Key   string `json:"key"`
}

type ApiHTTPRoute struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func init() {
	flag.Parse()
}

func main() {
	if *version {
		fmt.Printf("O(lya-lya) server build: %s\n", buildName)
		os.Exit(0)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	router := pat.New()

	router.Post("/create", handlerDatabaseCreate)
	router.Get("/list", handlerDatabaseList)
	router.Delete("/destroy", handlerInstanceDestroy)

	router.Post("/in/{instance}/arr/set", handlerInstanceSetArray)
	router.Post("/in/{instance}/arr/el/add", handlerInstanceArrAdd)
	router.Post("/in/{instance}/arr/el/set", handlerInstanceArrSet)
	router.Get("/in/{instance}/arr/el/get", handlerInstanceArrGet)
	router.Delete("/in/{instance}/arr/el/del", handlerInstanceArrDel)

	router.Post("/in/{instance}/hash/set", handlerInstanceSetHash)
	router.Post("/in/{instance}/hash/el/set", handlerInstanceHashSet)
	router.Get("/in/{instance}/hash/el/get", handlerInstanceHashGet)
	router.Delete("/in/{instance}/hash/el/del", handlerInstanceHashDel)

	router.Post("/in/{instance}/ttl/set", handlerInstanceTTLSet)
	router.Delete("/in/{instance}/ttl/del", handlerInstanceTTLDel)

	router.Post("/in/{instance}/set", handlerInstanceSetString)
	router.Get("/in/{instance}/get/{name}", handlerInstanceGet)
	router.Delete("/in/{instance}/delete/{name}", handlerInstanceDel)

	router.Get("/in/{instance}", handlerInstanceInfo)

	router.Get("/", handlerNotFount)
	http.Handle("/", router)

	go func() {
		_ = <-sigs
		log.Printf("Server stops")
		os.Exit(0)
	}()

	log.Printf("Server %s (build %s) listening on %s", ServerName, buildName, *httpAddress)

	log.Fatal(http.ListenAndServe(*httpAddress, nil))
}

func handlerInstanceTTLSet(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	res := new(RequestJsonTTL)
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	instanceName := r.URL.Query().Get(":instance")
	instance, err := Db.Get(instanceName)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	err = instance.SetTTL(res.Name, res.Expire)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	handlerJsonOk(w)
}

func handlerInstanceTTLDel(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	res := new(RequestJsonTTL)
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	instanceName := r.URL.Query().Get(":instance")
	instance, err := Db.Get(instanceName)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	instance.DelTTL(res.Name)

	handlerJsonOk(w)
}

func handlerNotFount(w http.ResponseWriter, r *http.Request) {
	log.Println("Not Found")
	handlerJson(w, http.StatusNotFound, &RespondJson{
		"status": "ERROR",
		"path":   r.URL.Path,
	})
}
func handlerInstanceArrAdd(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	res := new(RequestJsonArrayItem)
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	instanceName := r.URL.Query().Get(":instance")
	instance, err := Db.Get(instanceName)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	err = instance.ArrAdd(res.Name, res.Value)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	handlerJsonOk(w)
}

func handlerInstanceArrGet(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	res := new(RequestJsonArrayItem)
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	instanceName := r.URL.Query().Get(":instance")
	instance, err := Db.Get(instanceName)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	value, err := instance.ArrGet(res.Name, res.Index)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	handlerJson(w, http.StatusOK, &RespondJson{
		"status": "OK",
		"value":  value,
	})
}

func handlerInstanceArrDel(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	res := new(RequestJsonArrayItem)
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	instanceName := r.URL.Query().Get(":instance")
	instance, err := Db.Get(instanceName)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	err = instance.ArrDel(res.Name, res.Index)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	handlerJsonOk(w)
}

func handlerInstanceArrSet(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	res := new(RequestJsonArrayItem)
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	instanceName := r.URL.Query().Get(":instance")
	instance, err := Db.Get(instanceName)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	instance.ArrSet(res.Name, res.Index, res.Value)
	handlerJsonOk(w)
}

func handlerInstanceHashSet(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	res := new(RequestJsonHashItem)
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	instanceName := r.URL.Query().Get(":instance")
	instance, err := Db.Get(instanceName)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	err = instance.HashSet(res.Name, res.Key, res.Value)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	handlerJsonOk(w)
}

func handlerInstanceHashDel(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	res := new(RequestJsonHashItem)
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	instanceName := r.URL.Query().Get(":instance")
	instance, err := Db.Get(instanceName)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	err = instance.HashDel(res.Name, res.Key)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	handlerJsonOk(w)
}

func handlerInstanceHashGet(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	res := new(RequestJsonHashItem)
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	instanceName := r.URL.Query().Get(":instance")
	instance, err := Db.Get(instanceName)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	value, err := instance.HashGet(res.Name, res.Key)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	handlerJson(w, http.StatusOK, &RespondJson{
		"status": "OK",
		"value":  value,
	})
}

func handlerDatabaseCreate(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{}
	err := json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()

	if err != nil {
		handlerJsonError(w, err)
		return
	}

	err = Db.Create(data["name"])
	if err != nil {
		handlerJsonError(w, err)
		return
	}
	handlerJsonOk(w)
}

func handlerDatabaseList(w http.ResponseWriter, r *http.Request) {
	handlerJson(w, http.StatusOK, &RespondJson{
		"status": "OK",
		"count":  Db.Len(),
		"names":  Db.Keys(),
	})
}

func handlerInstanceDestroy(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{}
	err := json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()

	if err != nil {
		handlerJsonError(w, err)
		return
	}

	Db.Delete(data["name"])
	handlerJsonOk(w)
}

func handlerJson(w http.ResponseWriter, code int, object interface{}) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.Header().Set("Server", "OLL 0.0.1")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(object)
}

func handlerJsonError(w http.ResponseWriter, err error) {
	log.Println(err)
	handlerJson(w, http.StatusInternalServerError, &RespondJson{
		"status": "ERROR",
		"error":  err.Error(),
	})
}

func handlerJsonOk(w http.ResponseWriter) {
	handlerJson(w, http.StatusOK, &RespondJson{
		"status": "OK",
	})
}

func handlerInstanceInfo(w http.ResponseWriter, r *http.Request) {
	instanceName := r.URL.Query().Get(":instance")
	instance, err := Db.Get(instanceName)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	handlerJson(w, http.StatusOK, &RespondJson{
		"instance": instanceName,
		"length":   instance.Len(),
		"names":    instance.Keys(),
	})
}

func handlerInstanceGet(w http.ResponseWriter, r *http.Request) {
	instanceName := r.URL.Query().Get(":instance")
	instance, err := Db.Get(instanceName)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	name := r.URL.Query().Get(":name")
	value, err := instance.Get(name)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	handlerJson(w, http.StatusOK, &RespondJson{
		"value": value,
	})
}

func handlerInstanceDel(w http.ResponseWriter, r *http.Request) {
	instanceName := r.URL.Query().Get(":instance")
	instance, err := Db.Get(instanceName)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	name := r.URL.Query().Get(":name")
	instance.Del(name)

	handlerJsonOk(w)
}

func handlerInstanceSetString(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	res := new(RequestJsonString)
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	instanceName := r.URL.Query().Get(":instance")
	instance, err := Db.Get(instanceName)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	err = instance.Set(res.Name, res.Value)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	if res.Expire > 0 {
		err = instance.SetTTL(res.Name, res.Expire)
		if err != nil {
			handlerJsonError(w, err)
			return
		}
	}

	handlerJsonOk(w)
}

func handlerInstanceSetArray(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	res := new(RequestJsonArray)
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	instanceName := r.URL.Query().Get(":instance")
	instance, err := Db.Get(instanceName)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	err = instance.Set(res.Name, res.Value)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	if res.Expire > 0 {
		err = instance.SetTTL(res.Name, res.Expire)
		if err != nil {
			handlerJsonError(w, err)
			return
		}
	}

	handlerJsonOk(w)
}

func handlerInstanceSetHash(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	res := new(RequestJsonHash)
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	instanceName := r.URL.Query().Get(":instance")
	instance, err := Db.Get(instanceName)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	err = instance.Set(res.Name, res.Value)
	if err != nil {
		handlerJsonError(w, err)
		return
	}

	if res.Expire > 0 {
		err = instance.SetTTL(res.Name, res.Expire)
		if err != nil {
			handlerJsonError(w, err)
			return
		}
	}

	handlerJsonOk(w)
}
