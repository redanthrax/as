package services

import (
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azqueue"
	"github.com/redanthrax/as/api/internal/repository"
	"github.com/redanthrax/as/api/model"
)

type Pokemon interface {
  GetPokemon() ([]model.Pokemon, error)
  SyncPokemon() error
  GetPokemonQueue() (azqueue.PeekMessagesResponse, error)
}

type Services struct {
  Pokemon
}

func NewServices(repo *repository.Repository) *Services {
  return &Services{
    Pokemon: NewPokemonService(repo),
  }
}
