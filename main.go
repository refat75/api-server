package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	bookHandler := BookHandler{}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from book-server")
	})
	r.Post("/login", LoginHandler)
	r.Post("/logout", LogoutHandler)

	r.Route("/books", func(r chi.Router) {
		r.Use(JWTAuthMiddleware) // Apply middleware to all routes in this group
		r.Get("/", bookHandler.ListBooks)
		r.Post("/", bookHandler.CreateBook)
		r.Get("/{id}", bookHandler.GetBooks)
		r.Put("/{id}", bookHandler.UpdateBook)
		r.Delete("/{id}", bookHandler.DeleteBook)
	})

	err := http.ListenAndServe(":8081", r)
	if err != nil {
		fmt.Println("Could not start server", err)
	}
}
