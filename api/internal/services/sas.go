package services

import (
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azqueue"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azqueue/sas"
)

type SASService struct {
  Connection string
}

func NewSASService(connection string) *SASService {
  return &SASService{
    Connection: connection,
  }
}

func (s *SASService) GetSAS() (string, error) {
  //use the storage account stuff to generate an sas token
  client, err := azqueue.NewServiceClientFromConnectionString(s.Connection, nil)
  if err != nil {
    return "", err
  }

  resources := sas.AccountResourceTypes {Service: true}
  perms := sas.AccountPermissions {Read: true, Write: true, Delete: true}
  expiry := time.Now().AddDate(0, 0, 1)
  url, err := client.GetSASURL(resources, perms, expiry, nil)
  if err != nil {
    return "", err
  }

  return url, nil
}
