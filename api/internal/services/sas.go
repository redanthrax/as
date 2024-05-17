package services

type SASService struct {}

func NewSASService() *SASService {
  return &SASService{}
}

func (s *SASService) GetSAS() (string, error) {
  //use the storage account stuff to generate an sas token
  return "iamasastoken", nil
}
