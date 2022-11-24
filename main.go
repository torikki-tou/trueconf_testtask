package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/torikki-tou/trueconf_testtask/config"
	handlerV1 "github.com/torikki-tou/trueconf_testtask/handler/v1"
	"github.com/torikki-tou/trueconf_testtask/repo"
	"github.com/torikki-tou/trueconf_testtask/service"
)
var (
	userRepo 	repo.UserRepositiry 	= repo.NewUserRepository(config.Filename)
	userService service.UserService 	= service.NewUserService(userRepo)
	userHandler handlerV1.UserHandler 	= handlerV1.NewUserHandler(userService)
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/", userHandler.GetUsers)
				r.Post("/", userHandler.CreateUser)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", userHandler.GetUserByID)
					r.Patch("/", userHandler.UpdateUser)
					r.Delete("/", userHandler.DeleteUser)
				})
			})
		})
	})

	http.ListenAndServe(":3333", r)
}