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

type PokemonAzStorage struct {
  client *aztables.ServiceClient
  table *aztables.Client
  queue *azqueue.QueueClient
}

func NewPokemonAzStorage(db *aztables.ServiceClient, q *azqueue.ServiceClient) *PokemonAzStorage {
  //initialize and return
  db.CreateTable(context.Background(), "Pokemon", nil)
  q.CreateQueue(context.Background(), "pokemon", nil)
  return &PokemonAzStorage {
    client: db,
    table: db.NewClient("Pokemon"),
    queue: q.NewQueueClient("pokemon"),
  }
}

func (p *PokemonAzStorage) GetPokemon() ([]model.Pokemon, error) {
  pager := p.table.NewListEntitiesPager(nil)
  pokemon := make([]model.Pokemon, 0)
  for pager.More() {
    resp, err := pager.NextPage(context.Background())
    if err != nil {
      return nil, err
    }

    for _, entity := range resp.Entities {
      var poke model.Pokemon
      err = json.Unmarshal(entity, &poke)
      if err != nil {
        return nil, err
      }

      pokemon = append(pokemon, poke)
    }
  }

  return pokemon, nil
}

func(p *PokemonAzStorage) SyncPokemon() error {
  //storage queue requires a base64 encoded message
  msg := base64.StdEncoding.EncodeToString([]byte("hello"))
  resp, err := p.queue.EnqueueMessage(context.Background(), msg, nil)
  log.Info().Any("resp", resp).Msg("")
  if err != nil {
    log.Error().Err(err).Msg("")
    return err
  }

  return nil
}

func (p *PokemonAzStorage) GetPokemonQueue() (azqueue.PeekMessagesResponse, error) {
  msgs, err := p.queue.PeekMessages(context.Background(), nil)
  return msgs, err
}
