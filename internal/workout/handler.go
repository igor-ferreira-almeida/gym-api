package workout

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

type (
	serviceCreator interface {
		Add(ctx context.Context, w Workout) (Workout, error)
		Get(ctx context.Context, id string) (Workout, error)
	}
	service interface {
		serviceCreator
	}
)

type Handler struct {
	service service
}

func NewHandler(s service) Handler {
	return Handler{
		service: s,
	}
}

func (h Handler) Add(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("invalid body"))
	}
	var workout Workout
	json.Unmarshal(body, &workout)

	exercise, err := h.service.Add(r.Context(), workout)
	if err != nil {
		w.Write([]byte("service error"))
	}

	respBody, err := json.Marshal(exercise)
	if err != nil {
		w.Write([]byte("error in response"))
	}
	w.Write(respBody)
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "date")

	workout, err := h.service.Get(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("service error"))
	}
	resBody, err := json.Marshal(workout)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error in response"))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}
