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
type UrlCreationRequest struct {
	LongUrl string `json:"long_url"`
	UserId  string `json:"user_id"`
}

// func CreateShortUrl(w http.ResponseWriter, r *http.Request) {
// 	var creationRequest UrlCreationRequest
// 	if err :=
// }

func LandingPage(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "pkglication/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("halo"))
	fmt.Fprintf(w, "halow")
}

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

func HandleShortUrlRedirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortUrl := vars["shortUrl"]
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	http.Redirect(w, r, initialUrl, http.StatusSeeOther)
}
