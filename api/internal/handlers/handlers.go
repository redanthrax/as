package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/redanthrax/as/api/internal/services"
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
    r.Route("/pokemon", func(r chi.Router) {
      r.Get("/", h.GetPokemon)
      r.Get("/sync", h.SyncPokemon)
      r.Get("/queue", h.GetPokemonQueue)
    })
  })

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
