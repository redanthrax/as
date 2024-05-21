package handlers

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func (h *Handler) QueueTrigger(w http.ResponseWriter, r *http.Request) {
  log.Info().Msg("QueueTrigger")
  h.services.FetchPokemon()
  w.WriteHeader(http.StatusOK)
}
