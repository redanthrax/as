package repository

import (
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azqueue"
	"github.com/rs/zerolog/log"
)

type Config struct {
  StorageAccount string
}

func NewDB(config Config) (*aztables.ServiceClient, *azqueue.ServiceClient, error) {
  if config.StorageAccount == "" {
    log.Fatal().Msg("AZWebJobsStorage is not set")
  }

  client, err := aztables.NewServiceClientFromConnectionString(config.StorageAccount, nil)
  if err != nil {
    return nil, nil, err
  }

  queue, err := azqueue.NewServiceClientFromConnectionString(config.StorageAccount, nil)
  if err != nil {
    return nil, nil, err
  }

  return client, queue, nil
}
