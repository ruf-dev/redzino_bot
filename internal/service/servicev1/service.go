package servicev1

import (
	"github.com/ruf-dev/redzino_bot/internal/service"
	"github.com/ruf-dev/redzino_bot/internal/storage"
)

type Service struct {
	us *UserService
}

func NewService(dataStorage storage.Data) *Service {
	return &Service{
		us: NewUserService(dataStorage),
	}
}

func (s *Service) UserService() service.UserService {
	return s.us
}
