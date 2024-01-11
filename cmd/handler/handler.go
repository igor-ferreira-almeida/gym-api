package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/gym-api/internal/exercise"
	"github.com/gym-api/internal/workout"
	"log"
	"net/http"
)

func HelloWorldHandler(r *chi.Mux) {
	r.Get("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
}

func ExerciseHandler(r *chi.Mux) {
	h, err := exercise.HandlerInject()
	if err != nil {
		log.Fatalf("inject handler error: %v", err)
	}

	r.Post("/exercises", h.Add)
}

func WorkoutHandler(r *chi.Mux) {
	h, err := workout.HandlerInject()
	if err != nil {
		log.Fatalf("inject handler error: %v", err)
	}

	r.Post("/workouts", h.Add)
	r.Get("/workouts/{date}", h.Get)
}
