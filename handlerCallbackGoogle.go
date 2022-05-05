package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"surveyGenerator/web"
	"time"
)

type CallbackContent struct {
	ID             string
	Email          string
	Verified_email bool
	Picture        string
}

var callbackContent = CallbackContent{
	ID:             "",
	Email:          "",
	Verified_email: false,
	Picture:        "",
}

var authorizedEmails = []string{
	"benoit.frontzak@gmail.com",
	"carynfang@gmail.com",
}

var callbackGoogle = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/callback") {
		http.Error(w, "404 not found weird", http.StatusNotFound)
		return
	}

	switch r.Method {

	case "GET":
		getCallbackGoogle(w, r)

	default:
		log.Println("Sorry, only GET method is supported.")
	}

})

func getCallbackGoogle(w http.ResponseWriter, r *http.Request) {
	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	json.Unmarshal(content, &callbackContent)

	for _, authorizedEmail := range authorizedEmails {
		if callbackContent.Email == authorizedEmail {
			expiresDate := time.Now().Add(time.Minute * 60) // 60 min
			avatar := callbackContent.Picture
			login := strings.Split(callbackContent.Email, "@")[0]

			jwt, err := web.GenerateJWT(login, avatar, expiresDate)
			if err != nil {
				log.Println(err)
			}
			t := time.Now()

			// save token to cookie
			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   jwt,
				Expires: expiresDate,
			})
			log.Printf("session successfully saved in secured cookie for user %s at %s \n", callbackContent.Email, t.Format("2006-01-02 15:04:05"))
			// authorized
			http.Redirect(w, r, "/home", http.StatusMovedPermanently)
		}
	}
	// not authorized
	// http.Redirect(w, r, "/", http.StatusMovedPermanently)
	if err := tmpl.ExecuteTemplate(w, "unauthorized", nil); err != nil {
		log.Printf("%v", err)
	}
}

func getUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}
