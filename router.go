package main

import "net/http"

// Create handlers
func multiplexer() http.Handler {
	mux := http.NewServeMux()

	// File server
	staticDir := "/assets/"
	fs := http.FileServer(http.Dir("." + staticDir))
	mux.Handle(staticDir, http.StripPrefix(staticDir, fs))

	mux.Handle("/", home)                            // [GET] Home
	mux.Handle("/survey/generator", surveyGenerator) // [GET|POST] Generate Survey form
	mux.Handle("/survey/research/", surveyForm)      // [GET|POST] Survey form
	mux.Handle("/survey/feedback/", surveyFeedback)  // [GET] Survey feedback
	mux.Handle("/survey/report/", surveyReport)      // [GET] Survey report

	return mux
}
