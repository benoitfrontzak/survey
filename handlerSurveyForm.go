package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var surveyForm = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/survey/research/") {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {

	case "GET":
		getSurvey(w, r)

	case "POST":
		postSurvey(w, r)

	default:
		log.Println("Sorry, only GET method is supported.")
	}

})

func getSurvey(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	sURL := strings.Split(r.URL.Path, "/")
	sid := sURL[3]
	var d, err = getQuestionnaire(sid)
	if err != nil {
		log.Println(err)
	}

	// template data
	td := TemplateData{
		Surveys: getTemplateSurveys(),
		Data:    d,
	}

	log.Printf("home page successfully query at %s", t.Format("2006-01-02 15:04:05"))

	if err := tmpl.ExecuteTemplate(w, "survey", td); err != nil {
		log.Printf("%v", err)
	}
}

func postSurvey(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	sURL := strings.Split(r.URL.Path, "/")
	// Survey's ID
	sid := sURL[3]

	// Collect evaluation
	e, _ := collectEvaluation(r, sid)

	rowID, err := createDocument("evaluations", e)
	if err != nil {
		log.Println(err)
	}

	log.Printf("survey %s successfully posted at %s with id %s", sid, t.Format("2006-01-02 15:04:05"), rowID)
	http.Redirect(w, r, "/survey/feedback/"+rowID.(primitive.ObjectID).Hex(), http.StatusMovedPermanently)
}
