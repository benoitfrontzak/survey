package web

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	cookieName = "connectual"
	authIss    = "connectual-issuer"
	iv         = "dsay@x}(~q9kZE;w" // initialization vector
)

var jwtKey = []byte(iv)

// Custom claims for JSON Web Token
// type StandardClaims struct {
// 	Audience  string `json:"aud,omitempty"`
// 	ExpiresAt int64  `json:"exp,omitempty"`
// 	Id        string `json:"jti,omitempty"`
// 	IssuedAt  int64  `json:"iat,omitempty"`
// 	Issuer    string `json:"iss,omitempty"`
// 	NotBefore int64  `json:"nbf,omitempty"`
// 	Subject   string `json:"sub,omitempty"`
// }
type CustomClaims struct {
	Login      string `json:"login"`
	Avatar     string `json:"avatar"`
	Authorized bool   `json:"authorized"`
	jwt.StandardClaims
}

// Generate JSON Web Token
// A token is made of three parts, separated by .'s.
// The first part is called the header.
// It contains the necessary information for verifying the last part, the signature.
// The first two parts are JSON objects, that have been base64url encoded.
// The part in the middle is the interesting bit. It's called the Claims and contains the actual stuff you care about.
// Refer to RFC 7519 for information about reserved keys and the proper way to add your own.
// The last part is the signature, encoded the same way.

func GenerateJWT(login, avatar string, exp time.Time) (tk string, err error) {
	var (
		ck = []byte(iv)                      // cipher key
		t  = jwt.New(jwt.SigningMethodHS256) // token
		c  = t.Claims.(jwt.MapClaims)        // claims
		e  = exp.Unix()                      // claim's expiration time
	)

	c["login"] = login     // store authenticated login
	c["authorized"] = true // authorized used by middleware
	c["avatar"] = avatar   // guest|admin
	c["iss"] = authIss     // issuer used by middleware
	c["exp"] = e           // jwt expires time

	tk, err = t.SignedString(ck)
	if err != nil {
		return
	}

	return
}

func ReturnClaims(tk, iv string) (c *CustomClaims, err error) {
	c = &CustomClaims{}
	tkn, err := jwt.ParseWithClaims(tk, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(iv), nil
	})
	if err != nil {
		log.Println("claims returns")
		log.Println(err)
		return nil, err
	}
	if !tkn.Valid {
		err1 := errors.New("token is not valid")
		return nil, err1
	}
	return
}

func ExtractToken(r *http.Request, w http.ResponseWriter) (c *CustomClaims, err error) {
	// reqToken := r.Header.Get("Authorization")
	// splitToken := strings.Split(reqToken, "Bearer ")
	// reqToken = splitToken[len(splitToken)-1]
	// We can obtain the session token from the requests cookies, which come with every request
	cc, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the JWT string from the cookie
	tknStr := cc.Value

	// Initialize a new instance of `Claims`
	claims := &CustomClaims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	c, err = ReturnClaims(tknStr, iv)
	return
}
