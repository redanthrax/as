package repository

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
	"github.com/redanthrax/as/api/model"
)

type PokemonAzStorage struct {
  client *aztables.ServiceClient
  table *aztables.Client
}

func NewPokemonAzStorage(db *aztables.ServiceClient) *PokemonAzStorage {
  //initialize and return
  db.CreateTable(context.Background(), "Pokemon", nil)
  return &PokemonAzStorage {
    client: db,
    table: db.NewClient("Pokemon"),
  }
}

func (p *PokemonAzStorage) GetPokemon() ([]model.Pokemon, error) {
  options := &aztables.ListEntitiesOptions{}
  pager := p.table.NewListEntitiesPager(options)
  pokemon := make([]model.Pokemon, 0)
  for pager.More() {
    resp, err := pager.NextPage(context.Background())
    if err != nil {
      return nil, err
    }

    for _, entity := range resp.Entities {
      var poke model.Pokemon
      err = json.Unmarshal(entity, &poke)
      if err != nil {
        return nil, err
      }

      pokemon = append(pokemon, poke)
    }
  }

  return pokemon, nil
}
