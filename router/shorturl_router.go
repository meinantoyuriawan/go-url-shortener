package router

import (
	"github.com/gorilla/mux"
	"github.com/meinantoyuriawan/go-url-shortener/controller"
)

var RegisterShortUrlRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controller.LandingPage).Methods("GET")
	router.HandleFunc("/create-short-url", controller.CreateShortUrl).Methods("POST")
	router.HandleFunc("/go-to/{shortUrl}", controller.HandleShortUrlRedirect).Methods("GET")
}
