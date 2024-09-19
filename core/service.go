package core

type Config struct {
	timeout    int
	PrivateKey string
	PublicKey  string
}

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Post() {
}
