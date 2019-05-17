package router

import (
	"encoding/json"
	"log"
	"net/http"
)

type (
	HandleFunc func(http.ResponseWriter, *http.Request) map[string]interface{}
)

func (fn HandleFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")
	data := fn(w, r)
	w.WriteHeader(http.StatusBadRequest)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println(err)
		return
	}

}
