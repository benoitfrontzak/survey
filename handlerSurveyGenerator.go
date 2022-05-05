package main

import (
	"log"
	"net/http"
	"surveyGenerator/web"
	"time"
)

var surveyGenerator = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/survey/generator" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {

	case "GET":
		getGenerator(w, r)

	case "POST":
		postGenerator(w, r)

	default:
		log.Println("Sorry, only GET & POST method are supported.")
	}

})

func getGenerator(w http.ResponseWriter, r *http.Request) {
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

	log.Printf("survey generator page successfully query at %s", t.Format("2006-01-02 15:04:05"))

	if err := tmpl.ExecuteTemplate(w, "surveyGenerator", td); err != nil {
		log.Printf("%v", err)
	}
}

func postGenerator(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	q, _ := collectQuestionnaire(r)
	rowID, err := createDocument("questionnaires", q)
	if err != nil {
		log.Println(err)
	}

	log.Printf("survey %s successfully generated at %s", rowID, t.Format("2006-01-02 15:04:05"))
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
