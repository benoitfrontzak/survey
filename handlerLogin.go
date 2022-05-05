package main

import (
	"log"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)

func init() {
	googleOauthConfig = &oauth2.Config{

		// RedirectURL:  "http://localhost:8080/callback",
		RedirectURL:  "https://hrtools.herokuapp.com/callback",
		ClientID:     "11547521023.apps.googleusercontent.com",
		ClientSecret: "Qf_vPolsVLSQRitulNOUspAF",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

var login = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {

	case "GET":
		getLogin(w, r)

	case "POST":
		postLogin(w, r)

	default:
		log.Println("Sorry, only GET & POST method are supported.")
	}

})

func getLogin(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	log.Printf("login page successfully query at %s", t.Format("2006-01-02 15:04:05"))
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}

func postLogin(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	q, _ := collectQuestionnaire(r)
	rowID, err := createDocument("questionnaires", q)
	if err != nil {
		log.Println(err)
	}

	log.Printf("survey %s successfully generated at %s", rowID, t.Format("2006-01-02 15:04:05"))
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
