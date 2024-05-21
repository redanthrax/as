package services

import (
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azqueue"
	"github.com/redanthrax/as/api/internal/repository"
	"github.com/redanthrax/as/api/model"
)

type PokemonService struct {
  repo repository.Pokemon
}

func NewPokemonService(repo repository.Pokemon) *PokemonService {
  return &PokemonService{repo: repo}
}

func (s *PokemonService) GetPokemon() ([]model.Pokemon, error) {
  poke, err := s.repo.GetPokemon()
  return poke, err
}

func (s *PokemonService) SyncPokemon() error {
  err := s.repo.SyncPokemon()
  return err
}

func (s *PokemonService) GetPokemonQueue() (azqueue.PeekMessagesResponse, error) {
  msgs, err := s.repo.GetPokemonQueue()
  return msgs, err
}
