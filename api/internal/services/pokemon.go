package services

import (
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
