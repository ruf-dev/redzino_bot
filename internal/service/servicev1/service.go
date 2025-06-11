package servicev1

import (
	"github.com/ruf-dev/redzino_bot/internal/service"
	"github.com/ruf-dev/redzino_bot/internal/storage"
	"github.com/ruf-dev/redzino_bot/internal/storage/tx_manager"
)

type Service struct {
	userService            *UserService
	motivationService      *MotivationService
	chatService            *ChatService
	dailyActivitiesService *DailyActivitiesService
}

func (s *Service) MotivationService() service.MotivationService {
	return s.motivationService
}

func NewService(dataStorage storage.Data, txManager *tx_manager.TxManager) *Service {
	return &Service{
		userService:            NewUserService(dataStorage, txManager),
		motivationService:      NewMotivationService(dataStorage, txManager),
		chatService:            NewChatService(dataStorage),
		dailyActivitiesService: NewDailyActivitiesService(txManager, dataStorage),
	}
}

func (s *Service) UserService() service.UserService {
	return s.userService
}

func (s *Service) ChatService() service.ChatService {
	return s.chatService
}

func (s *Service) DailyActivitiesService() service.DailyActivitiesService {
	return s.dailyActivitiesService
}
