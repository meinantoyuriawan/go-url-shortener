package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/meinantoyuriawan/go-url-shortener/router"
	"github.com/meinantoyuriawan/go-url-shortener/store"

	"github.com/go-openapi/runtime/middleware"
)

func main() {
	r := mux.NewRouter()
	router.RegisterShortUrlRoutes(r)
	store.InitializeStore()
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	r.Handle("/docs", sh)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}
