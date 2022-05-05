package main

import (
	"log"
	"net/http"
	"strings"
	"surveyGenerator/web"
	"time"
)

var surveyPDF = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/survey/download/") {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {

	case "GET":
		PDFPage(w, r)

	default:
		log.Println("Sorry, only GET method is supported.")
	}

})

func PDFPage(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	// Retrieve the claims from the request (bearer)
	claims, err := web.ExtractToken(r, w)
	if err != nil {
		log.Println(err)
	}

	sURL := strings.Split(r.URL.Path, "/")
	// Survey's ID
	fid := sURL[3]
	var d, _ = getQuestionnaire(fid)

	// template data
	td := TemplateData{
		Login:   claims.Login,
		Avatar:  claims.Avatar,
		Surveys: getTemplateSurveys(),
		Data:    d,
	}

	log.Printf("home page successfully query at %s", t.Format("2006-01-02 15:04:05"))

	if err := tmpl.ExecuteTemplate(w, "surveyPDF", td); err != nil {
		log.Printf("%v", err)
	}
}
