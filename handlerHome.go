package main

import (
	"log"
	"net/http"
	"surveyGenerator/web"
	"time"
)

var home = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/home" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {

	case "GET":
		homepage(w, r)

	default:
		log.Println("Sorry, only GET method is supported.")
	}

})

func homepage(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	// Retrieve the claims from the request (bearer)
	claims, err := web.ExtractToken(r, w)
	if err != nil {
		log.Println(err)
	}

	// template data
	td := TemplateData{
		Login:   claims.Login,
		Avatar:  claims.Avatar,
		Surveys: getTemplateSurveys(),
		Data:    Questionnaire{},
	}

	log.Printf("home page successfully query at %s", t.Format("2006-01-02 15:04:05"))

	if err := tmpl.ExecuteTemplate(w, "home", td); err != nil {
		log.Printf("%v", err)
	}
}
