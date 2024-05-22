package services

import (
	"github.com/mtslzr/pokeapi-go"
	"github.com/redanthrax/as/worker/internal/repository"
	"github.com/redanthrax/as/worker/model"
	"github.com/rs/zerolog/log"
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
  //sync up the db
  log.Info().Msg("Syncing pokemon....")
  poke, err := pokeapi.Resource("pokemon", 0, 2000)
  if err != nil {
    return err
  }
  
  for _, p := range poke.Results {
    pm := model.Pokemon {
      Name: p.Name,
    }

    log.Info().Str("pokemon", p.Name).Msg("adding pokemon")
    s.repo.AddPokemon(pm)
  }

  log.Info().Msg("Pokemon sync complete")
  return nil
}
