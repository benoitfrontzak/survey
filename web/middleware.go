package web

import (
	"log"
	"net/http"
	"time"
)

// This function check if http request is authorized
// by retrieving the claims from the securecookie and
// check if claims.Authorized is true
// check if claims.Issuer is $authIss
// check if claims.ExpiresAt is within the last 10 min
// and renew it if needed (auto-renew)

func AuthMiddleware(next http.Handler) http.Handler {
	// Return Handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the claims from the request (bearer)
		claims, err := ExtractToken(r, w)
		if err != nil {
			log.Println(err)
		}

		// Define end of token (10 min before it ends)
		// so let's check 10 min away from now...
		now := time.Now().Add(time.Minute * 10).Unix()

		// 	// Check if claim is authorized and issued by $authIss
		if claims.Authorized && claims.Issuer == authIss {

			// Check if claim's expiration is within the last 5 min
			if now < claims.ExpiresAt {
				next.ServeHTTP(w, r)
			} else {
				// When claim's expiration is about to end
				expiresDate := time.Now().Add(time.Minute * 60) // 60 min
				claims.ExpiresAt = expiresDate.Unix()
				log.Printf("%s\n", "token successfully renewed")
				next.ServeHTTP(w, r)
			}

		} else {
			// The request is not authorized
			http.Redirect(w, r, "/unauthorized", http.StatusForbidden)
		}

	})

}
