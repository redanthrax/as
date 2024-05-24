
package repository

import (
	"context"
	"encoding/base64"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azqueue"
	"github.com/redanthrax/as/api/model"
	"github.com/rs/zerolog/log"
)

type Pokemon interface {
  GetPokemon() ([]model.Pokemon, error)
  SyncPokemon() error
  GetPokemonQueue() (azqueue.PeekMessagesResponse, error)
}

type Repository struct {
  Pokemon
}

func NewRepository(db *aztables.ServiceClient, q *azqueue.ServiceClient) *Repository {
  repo := &Repository{}
  repo.Pokemon = NewPokemonAzStorage(repo, db, q)
  return repo
}

func (r *Repository) SendCommand(command model.Command, queue *azqueue.QueueClient) error {
  js, err := json.Marshal(command)
  msg := base64.StdEncoding.EncodeToString(js)
  resp, err := queue.EnqueueMessage(context.Background(), msg, nil)
  log.Info().Any("resp", resp).Msg("")
  if err != nil {
    log.Error().Err(err)
    return err
  }

  return nil
}
