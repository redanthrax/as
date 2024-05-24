package repository

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azqueue"
	"github.com/rs/zerolog/log"
)

func (p *PokemonAzStorage) GetPokemonQ() (azqueue.PeekMessagesResponse, error) {
  msgs, err := p.queue.PeekMessages(context.Background(), nil)
  return msgs, err
}

func (p *PokemonAzStorage) SyncPokemonQ() error {
  log.Info().Msg("SyncPokemonQ")
  return nil
}
