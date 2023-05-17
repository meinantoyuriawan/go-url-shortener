package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/meinantoyuriawan/go-url-shortener/router"
	"github.com/meinantoyuriawan/go-url-shortener/store"
)

func main() {
	r := mux.NewRouter()
	router.RegisterShortUrlRoutes(r)
	store.InitializeStore()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
