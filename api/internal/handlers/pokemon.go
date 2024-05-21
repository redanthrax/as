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
