package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/redanthrax/as/worker/internal/services"
)

type Handler struct {
  services *services.Services
}

func NewHandler(services *services.Services) *Handler {
  return &Handler{services: services}
}

func (h *Handler) InitRoutes() http.Handler {
  r := chi.NewRouter()
  r.Use(middleware.Timeout(30 * time.Second))
  r.HandleFunc("/QueueTrigger", h.QueueTrigger)
  return r
}

func (h *Handler) HandleError(err error, w http.ResponseWriter) {
  http.Error(w, err.Error(), http.StatusInternalServerError)
}

func (h *Handler) HandleResponse(obj any, w http.ResponseWriter) {
	if err := json.NewEncoder(w).Encode(obj); err != nil {
    h.HandleError(err, w)
    return
	}
}
