package servicev1

import (
	"github.com/ruf-dev/redzino_bot/internal/service"
	"github.com/ruf-dev/redzino_bot/internal/storage"
	"github.com/ruf-dev/redzino_bot/internal/storage/tx_manager"
)

type Service struct {
	userService       *UserService
	motivationService *MotivationService
	chatService       *ChatService
}

func (s *Service) MotivationService() service.MotivationService {
	return s.motivationService
}

func NewService(dataStorage storage.Data, manager *tx_manager.TxManager) *Service {
	return &Service{
		userService:       NewUserService(dataStorage),
		motivationService: NewMotivationService(dataStorage, manager),
		chatService:       NewChatService(dataStorage),
	}
}

func (s *Service) UserService() service.UserService {
	return s.userService
}

func (s *Service) ChatService() service.ChatService {
	return s.chatService
}
