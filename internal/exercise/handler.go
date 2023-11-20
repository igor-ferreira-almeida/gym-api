package exercise

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type (
	serviceCreator interface {
		Add(ctx context.Context, e Exercise) (Exercise, error)
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
	var e Exercise
	json.Unmarshal(body, &e)

	//bigqueue.Send()

	exercise, err := h.service.Add(r.Context(), e)
	if err != nil {
		w.Write([]byte("service error"))
	}

	respBody, err := json.Marshal(exercise)
	if err != nil {
		w.Write([]byte("error in response"))
	}
	w.Write(respBody)
}

func (h Handler) Get(writer http.ResponseWriter, request *http.Request) {}
