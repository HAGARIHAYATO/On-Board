package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(CORS)
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/works", func(r chi.Router) {
			r.Get("/", GetWorks)
			r.Post("/", CreateWorks)
			r.Route("/{workID}", func(r chi.Router) {
				r.Get("/", GetWorkByID)
				r.Put("/", UpdateWorks)
				r.Put("/edit_item", UpdateWorkItem)
				r.Delete("/", DeleteWorks)
			})
		})
		r.Route("/users", func(r chi.Router) {
			r.Get("/", GetUsers)
			r.Route("/{userID}", func(r chi.Router) {
				r.Get("/", GetUserByID)
				r.Put("/", UpdateUser)
			})
		})
		r.Post("/login", RequestSession)
		r.Post("/signup", CreateUser)
		r.Get("/credential", GetPrivateInfo)
		r.Delete("/delete_account", DeleteUser)
	})
	http.ListenAndServe(":8080", r)
}
