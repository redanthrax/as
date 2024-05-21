package services

import (
	"time"

	//"github.com/mtslzr/pokeapi-go"
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

func (s *PokemonService) FetchPokemon() error {
  //sync up the db
  //poke := pokeapi.Resource("pokemon")
  log.Info().Msg("Fetching pokemon....")
  time.Sleep(time.Second * 10) 
  log.Info().Msg("Pokemon fetched.")
  return nil
}
