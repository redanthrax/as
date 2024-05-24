package repository

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azqueue"
)

type PokemonAzStorage struct {
  repo *Repository
  client *aztables.ServiceClient
  table *aztables.Client
  queue *azqueue.QueueClient
}

func NewPokemonAzStorage(repo *Repository, db *aztables.ServiceClient, q *azqueue.ServiceClient) *PokemonAzStorage {
  //initialize and return
  db.CreateTable(context.Background(), "Pokemon", nil)
  q.CreateQueue(context.Background(), "pokemon", nil)

  return &PokemonAzStorage {
    repo: repo,
    client: db,
    table: db.NewClient("Pokemon"),
    queue: q.NewQueueClient("pokemon"),
  }
}
