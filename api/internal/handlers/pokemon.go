package handlers

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func (h *Handler) GetPokemon(w http.ResponseWriter, r *http.Request) {
  log.Info().Msg("GetPokemon")
  poke, err := h.services.GetPokemon()
  if err != nil {
    h.HandleError(err, w)
  }

  h.HandleResponse(poke, w)
}

func (h *Handler) SyncPokemon(w http.ResponseWriter, r *http.Request) {
  log.Info().Msg("SyncPokemon")
  err := h.services.SyncPokemon()
  if err != nil {
    h.HandleError(err, w)
  }
}

func (h *Handler) GetPokemonQueue(w http.ResponseWriter, r *http.Request) {
  log.Info().Msg("GetPokemonQueue")
  msgs, err := h.services.GetPokemonQueue()
  if err != nil {
    h.HandleError(err, w)
  }

  h.HandleResponse(msgs, w)
}
