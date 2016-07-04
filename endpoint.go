package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamada")
	// Only handle POST requests, validate data
	if r.Method == "POST" {
		l := LogEntry{}
		json.NewDecoder(r.Body).Decode(&l)
		l.id = bson.NewObjectId()
		e := l.Persist()
		if e != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe("192.168.33.10:8080", nil)
}
