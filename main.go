package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message" : "get called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message" : "put called"}`))
}

func name(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"message" : "get %s called"}`, params["name"])))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message" : "post called"}`))
}

func postMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"postMessage %s called"}`, html.EscapeString(r.FormValue("message")))))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message" : "not found"}`))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/test/{name}", name).Methods(http.MethodGet)
	r.HandleFunc("/", put).Methods(http.MethodPut)
	r.HandleFunc("/", post).Methods(http.MethodPost)
	r.HandleFunc("/message", postMessage).Methods(http.MethodPost)
	r.HandleFunc("/", notFound)
	log.Println("Server started on: http://localhost:8080")
	r.HandleFunc("/index", Index)
	r.HandleFunc("/show", Show)
	r.HandleFunc("/new", New)
	r.HandleFunc("/edit", Edit)
	r.HandleFunc("/insert", Insert)
	r.HandleFunc("/update", Update)
	r.HandleFunc("/delete", Delete)
	log.Fatal(http.ListenAndServe(":8080", r))
}
