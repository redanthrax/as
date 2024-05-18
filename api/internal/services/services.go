package services

type SAS interface {
  GetSAS() (string, error)
}

type Services struct {
  SAS
}

func NewServices(connection string) *Services {
  return &Services{
    SAS: NewSASService(connection),
  }
}
