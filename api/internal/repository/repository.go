package repository

import (
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
	"github.com/redanthrax/as/api/model"
)

type Pokemon interface {
  GetPokemon() ([]model.Pokemon, error)
}

type Repository struct {
  Pokemon
}

func NewRepository(db *aztables.ServiceClient) *Repository {
  return &Repository{
    Pokemon: NewPokemonAzStorage(db),
  }
}
