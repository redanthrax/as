package repository

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
	"github.com/redanthrax/as/api/model"
)


func (p *PokemonAzStorage) GetPokemonT() ([]model.Pokemon, error) {
  pager := p.table.NewListEntitiesPager(nil)
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

func (p *PokemonAzStorage) AddPokemonT(pokemon model.Pokemon) error {
  properties := map[string]interface{}{
    "Name": pokemon.Name,
  }

  entity := aztables.EDMEntity {
    Entity: aztables.Entity {
      RowKey: pokemon.Name,
      PartitionKey: "pokemon",
    },
    Properties: properties,
  }

  marshalled, err := json.Marshal(entity)
  if err != nil {
    return err
  }

  _, err = p.table.AddEntity(context.TODO(), marshalled, nil)
  if err != nil {
    return err
  }

  return nil
}

func(p *PokemonAzStorage) SyncPokemonT() error {
  com := model.Command {
    Function: "SyncPokemon",
  }

  err := p.repo.SendCommand(com, p.queue)
  return err
}


