package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

var surveyFeedback = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/survey/feedback/") {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {

	case "GET":
		feedbackPage(w, r)

	default:
		log.Println("Sorry, only GET method is supported.")
	}

})

func feedbackPage(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	sURL := strings.Split(r.URL.Path, "/")
	// Survey's ID
	fid := sURL[3]
	d, _ := collectFeedback(fid)
	// template data
	td := TemplateData{
		Surveys: getTemplateSurveys(),
		Data:    d,
	}

	log.Printf("home page successfully query at %s", t.Format("2006-01-02 15:04:05"))

	if err := tmpl.ExecuteTemplate(w, "surveyFeedback", td); err != nil {
		log.Printf("%v", err)
	}
}
