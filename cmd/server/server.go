package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gym-api/cmd/handler"
	"github.com/rs/cors"
	"net/http"
)

func NewServer() {
	r := chi.NewRouter()

	// Use default options
	r.Use(cors.AllowAll().Handler)
	r.Use(middleware.Logger)

	handler.HelloWorldHandler(r)
	handler.ExerciseHandler(r)
	handler.WorkoutHandler(r)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
