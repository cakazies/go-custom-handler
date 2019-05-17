package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/local/go-custom-handler/router"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/index", router.HandleFunc(router.Index))
	r.Handle("/home", router.HandleFunc(router.Home))

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8181",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
