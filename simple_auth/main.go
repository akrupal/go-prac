package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var secretKey = []byte("internal_api_key")

type UsernameString string

var Username UsernameString = "Username"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if user.Username == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Username or password missing")
		return
	}
	// skipping password verification
	expTime := time.Now().Add(time.Minute * 5)
	claim := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(secretKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "Token",
		Value:   token,
		Expires: expTime,
	})

	//you could also set it as a header
	// w.Header().Set("Token", token)
}

func UserDataHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value(Username).(string)

	json.NewEncoder(w).Encode(fmt.Sprintf("Hello %s", username))
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		}
		token := cookie.Value

		//otherwise if header was used to pass the token
		// token:=r.Header.Get("Token")

		claim := &Claims{}
		tkn, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err != nil || !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
		}

		ctx := context.WithValue(r.Context(), Username, claim.Username)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// if some argument needs to be passed to the middleware
func NewMiddleware(keys string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// compare keys
			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/Login", LoginHandler)

	api := r.PathPrefix("/User").Subrouter()
	api.HandleFunc("/Data", UserDataHandler)
	api.Use(Middleware)

	http.ListenAndServe(":8080", r)
}
