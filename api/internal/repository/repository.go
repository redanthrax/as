package repository

import (
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azqueue"
	"github.com/redanthrax/as/api/model"
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
  return &Repository{
    Pokemon: NewPokemonAzStorage(db, q),
  }
}
