package servicev1

import (
	"github.com/ruf-dev/redzino_bot/internal/service"
	"github.com/ruf-dev/redzino_bot/internal/storage"
)

type Service struct {
	us *UserService
	ms *MotivationService
}

func (s *Service) MotivationService() service.MotivationService {
	return s.ms
}

func NewService(dataStorage storage.Data) *Service {
	return &Service{
		us: NewUserService(dataStorage),
		ms: NewMotivationService(dataStorage),
	}
}

func (s *Service) UserService() service.UserService {
	return s.us
}
