package main

import (
	"html/template"
	"log"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Define functions for templates
var fns = template.FuncMap{
	"inc": func(x int) int {
		return x + 1
	},
	"dec": func(x int) int {
		return x - 1
	},
	"sum": func(n, t int) int {
		return n + t
	},
	"atoi": func(s string) int {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Println(err)
		}
		return i
	},
	"percentage": func(n, t int) int {
		r := float64(n) / float64(t) * 100
		return int(r)
	},
	"rowID": func(x primitive.ObjectID) string {
		return string(x.Hex())
	},
	"mkSlice": func(args ...string) []string {
		return args
	},
	"isInclude": func(s string, a []string) bool {
		inc := false
		for _, e := range a {
			if e == s {
				inc = true
			}
		}
		return inc
	},
	"mkSliceFromSpace": func(s string) []string {
		slc := strings.Fields(s)
		return slc
	},
}

// Define templates folder
var tmpl = template.Must(template.New("base").Funcs(fns).ParseGlob("templates/*"))
