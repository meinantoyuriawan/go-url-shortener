package controller

import (
	// "encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/meinantoyuriawan/go-url-shortener/shortener"
	"github.com/meinantoyuriawan/go-url-shortener/store"
	"github.com/meinantoyuriawan/go-url-shortener/utils"
)

// Request model definition

// swagger:model UrlCreationRequest
type UrlCreationRequest struct {
	// name: LongUrl
	// in: string
	LongUrl string `json:"long_url"`
	// name: UserId
	// in: string
	UserId string `json:"user_id"`
}

// func CreateShortUrl(w http.ResponseWriter, r *http.Request) {
// 	var creationRequest UrlCreationRequest
// 	if err :=
// }

// swagger:route GET / UrlCreationRequest getLanding
// Get landing page
//
// security:
// - apiKey: []
// responses:
//
//	401: Error
//	200: StringHello
func LandingPage(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "pkglication/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("halo"))
	fmt.Fprintf(w, "halow")
}

// swagger:route POST /create-short-url UrlCreationRequest createUrl
// Create a short url
//
// security:
// - apiKey: []
// responses:
//
//	401: Error
//	200: short url
func CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	createShortUrl := &UrlCreationRequest{} //accessing the book struct
	utils.ParseBody(r, createShortUrl)
	shortUrl := shortener.GenerateShortLink(createShortUrl.LongUrl, createShortUrl.UserId)

	store.SaveUrlMapping(shortUrl, createShortUrl.LongUrl, createShortUrl.UserId)

	// res, _ := json.Marshal(createShortUrl)
	// w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(shortUrl))
}

// swagger:route GET /go-to/{shortUrl} UrlCreationRequest goToUrl
// redirecting to desired url (short)
//
// security:
// - apiKey: []
// responses:
//
//	401: Error
//	200: Success message
func HandleShortUrlRedirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortUrl := vars["shortUrl"]
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	http.Redirect(w, r, initialUrl, http.StatusSeeOther)
}
