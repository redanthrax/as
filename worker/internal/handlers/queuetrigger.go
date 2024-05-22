package handlers

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func (h *Handler) QueueTrigger(w http.ResponseWriter, r *http.Request) {
  log.Info().Msg("QueueTrigger")

  err := h.services.SyncPokemon()
  if err != nil {
    log.Err(err).Msg("")
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  w.WriteHeader(http.StatusOK)
  return
}
