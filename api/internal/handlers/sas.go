package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func (h *Handler) GetSAS(w http.ResponseWriter, r *http.Request) {
  log.Info().Msg("GETSAS")
  sas, err := h.services.SAS.GetSAS()
  if err != nil {
    return
  }

  if err := json.NewEncoder(w).Encode(sas); err != nil {
    return
  }
}
