package services

import (
  "github.com/mtslzr/pokeapi-go"
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
  poke, err := s.repo.GetPokemonT()
  return poke, err
}

func (s *PokemonService) SyncPokemon() error {
  err := s.repo.SyncPokemonT()
  return err
}

func (s *PokemonService) GetPokemonQueue() (azqueue.PeekMessagesResponse, error) {
  msgs, err := s.repo.GetPokemonQ()
  return msgs, err
}

func (s *PokemonService) PullPokemon() ([]model.Pokemon, error) {
  poke, err := pokeapi.Resource("pokemon", 0, 2000)
  if err != nil {
    return nil, err
  }
 
  pokemon := make([]model.Pokemon, 0)
  for _, p := range poke.Results {
    pm := model.Pokemon {
      Name: p.Name,
    }

    pokemon = append(pokemon, pm)
  }

  return pokemon, nil
}
