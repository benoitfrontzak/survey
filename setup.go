package main

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Customize prefix log
func setLog() {
	log.SetPrefix("[INFO]: ")
	log.SetFlags(0) // remove file:line and timestamps from log liness
}

// This function take 2 parameters:
// 	expected Variable to be found in .env
// 	expected Default value if not found
// Check if .env exist
// Check if variable set inside .env
// Returns either .env value or default value
func getEnv(value, defaut string) string {
	setLog()
	// Get env var
	if err := godotenv.Load(); err != nil {
		log.Printf("%s\n", "error loading .env file, set default value to "+defaut)
		return defaut
	}

	// Defined default value If not defined in env
	val := os.Getenv(value)

	if val == "" {

		val = ""
		log.Printf("%s %s %s%s%s\n", "error loading", value, ", set default value to `", defaut, "`")
		return val
	}

	return val
}

// gotEnv
func gotEnv() (exist bool) {
	if _, err := os.Stat(".env"); errors.Is(err, os.ErrNotExist) {
		exist = false
	} else {
		exist = true
	}
	return
}
