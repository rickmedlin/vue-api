package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// routes generates our routes and attaches them to handlers.
func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Post("/users/login", app.Login)
	mux.Post("/users/logout", app.Logout)
	mux.Post("/books", app.AllBooks)
	mux.Get("/books", app.AllBooks)
	mux.Get("/books/{slug}", app.OneBook)

	mux.Post("/validate-token", app.ValidateToken)

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(app.AuthTokenMiddleware)

		// admin user routes
		mux.Post("/users", app.AllUsers)
		mux.Post("/users/save", app.EditUser)
		mux.Post("/users/get/{id}", app.GetUser)
		mux.Post("/users/delete", app.DeleteUser)
		mux.Post("/log-user-out/{id}", app.LogUserOutAndSetInactive)

		// admin author routes
		mux.Post("/authors/all", app.AuthorsAll)
		mux.Post("/books/save", app.EditBook)

	})

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
