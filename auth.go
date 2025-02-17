package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

var secretKey = []byte("my-secret-key")

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if u.Username == "admin" && u.Password == "123456" {
		expirationTime := time.Now().Add(20 * time.Minute)
		tokenString, err := createToken(u.Username, expirationTime)
		if err != nil {
			http.Error(w, "Error creating token", http.StatusInternalServerError)
			return
		}

		//Store token in the cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    tokenString,
			Expires:  expirationTime,
			HttpOnly: true,
			Path:     "/",
		})

		fmt.Printf("Token generated successfull for user: %s\n", u.Username)
		w.WriteHeader(http.StatusOK)
		//fmt.Fprintf(w, tokenString)
		return
	}

	http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
}

func createToken(username string, expirationTime time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      expirationTime.Unix(),
	})

	return token.SignedString(secretKey)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Minute),
		HttpOnly: true,
		Path:     "/",
	})

	w.Write([]byte("Successfully logged out"))
	w.WriteHeader(http.StatusOK)
}

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tokenString := cookie.Value
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		username, _ := claims["username"].(string)

		if err != nil || !token.Valid || username != "admin" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
