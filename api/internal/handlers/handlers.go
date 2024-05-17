package handlers

import (
	"net/http"

	"github.com/redanthrax/as/api/internal/services"
	chi "github.com/go-chi/chi/v5"
)

type Handler struct {
  services *services.Services
}

func NewHandler(services *services.Services) *Handler {
  return &Handler{services: services}
}

func (h *Handler) InitRoutes() http.Handler {
  r := chi.NewRouter()
  r.Route("/api", func(r chi.Router) {
    r.Route("/getsas", func(r chi.Router) {
      r.Get("/", h.GetSAS)
    })
  })

  return r
}

func (h *Handler) HandleError(err error, w http.ResponseWriter) {
  http.Error(w, err.Error(), http.StatusInternalServerError)
}
