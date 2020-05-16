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
				r.Get("/information", GetInformation)
				r.Put("/", UpdateUser)
			})
		})
		r.Post("/login", RequestSession)
		r.Post("/signup", CreateUser)
		r.Post("/information", PostInformation)
		r.Get("/credential", GetPrivateInfo)
		r.Delete("/delete_account", DeleteUser)
		r.Post("/execute_account", ExecutedUser)
		r.Get("/works_ids", GetWorksIDs)
		r.Get("/users_ids", GetUsersIDs)
		r.Get("/skills", GetSkills)
		r.Get("/works_per_day", GetWorksPerDay)
	})
	http.ListenAndServe(":8080", r)
}
