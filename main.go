package main

import (
	"log"

	"surveyGenerator/web"
)

func main() {
	env := gotEnv()
	var httpPort string

	// If .env exist then we got a mongoDB conf
	if env {
		httpPort = getEnv("PORT", "8080")
		if pingDB := pingMongoDB(); pingDB != nil {
			log.Fatalf(pingDB.Error())
		}
	} else {
		httpPort = "8080"
	}

	// Start http server
	if web := web.Run(multiplexer(), httpPort); web != nil {
		log.Fatalf(web.Error())
	}
}
