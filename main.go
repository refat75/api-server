package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	r.Post("/login", LoginHandler)
	r.Post("/logout", LogoutHandler)

	r.Route("/api", func(r chi.Router) {
		r.Use(JWTAuthMiddleware) // Apply middleware to all routes in this group
		r.Get("/protected", ProtectedHandler)
		r.Get("/dashboard", DashboardHandler)
	})

	err := http.ListenAndServe(":8081", r)
	if err != nil {
		fmt.Println("Could not start server", err)
	}
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Welcome to the protected area")
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Welcome to the dashboard")
}
