package services

import (
	"github.com/redanthrax/as/worker/internal/repository"
	"github.com/redanthrax/as/worker/model"
)

type Pokemon interface {
  GetPokemon() ([]model.Pokemon, error)
  SyncPokemon() error
}

type Services struct {
  Pokemon
}

func NewServices(repo *repository.Repository) *Services {
  return &Services{
    Pokemon: NewPokemonService(repo),
  }
}
