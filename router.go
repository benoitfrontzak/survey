package main

import (
	"net/http"
	"surveyGenerator/web"
)

// Create handlers
func multiplexer() http.Handler {
	mux := http.NewServeMux()

	// File server
	staticDir := "/assets/"
	fs := http.FileServer(http.Dir("." + staticDir))
	mux.Handle(staticDir, http.StripPrefix(staticDir, fs))

	mux.Handle("/", login)                                               // [GET|POST] Login
	mux.Handle("/callback", callbackGoogle)                              // [GET] Login
	mux.Handle("/home", web.AuthMiddleware(home))                        // [GET] Home
	mux.Handle("/survey/generator", web.AuthMiddleware(surveyGenerator)) // [GET|POST] Generate Survey form
	mux.Handle("/survey/research/", surveyForm)                          // [GET|POST] Survey form
	mux.Handle("/survey/feedback/", surveyFeedback)                      // [GET] Survey feedback
	mux.Handle("/survey/report/", web.AuthMiddleware(surveyReport))      // [GET] Survey report
	mux.Handle("/survey/download/", web.AuthMiddleware(surveyPDF))       // [GET] Survey PDF

	return mux
}
