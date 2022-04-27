package main

import (
	"log"
	"strings"

	"surveyGenerator/web"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		log.Println(err)
	}

	htmlStr := `<html><body><h1 style="color:red;">This is an html
	 from pdf to test color<h1><img src="logo.png" alt="img" height="42" width="42"></img></body></html>`

	pdfg.AddPage(wkhtml.NewPageReader(strings.NewReader(htmlStr)))

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Println(err)
	}

	//Your Pdf Name
	err = pdfg.WriteFile("./checkthis.pdf")
	if err != nil {
		log.Println(err)
	}

	env := gotEnv()
	var httpPort string

	// If .env exist then we got a mongoDB conf
	if env {
		httpPort = getEnv("HTTPPORT", "8888")
		if pingDB := pingMongoDB(); pingDB != nil {
			log.Fatalf(pingDB.Error())
		}
	} else {
		httpPort = "8888"
	}

	// Start http server
	if web := web.Run(multiplexer(), httpPort); web != nil {
		log.Fatalf(web.Error())
	}
}
