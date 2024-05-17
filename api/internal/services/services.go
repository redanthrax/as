package services

type SAS interface {
  GetSAS() (string, error)
}

type Services struct {
  SAS
}

func NewServices() *Services {
  return &Services{
    SAS: NewSASService(),
  }
}
