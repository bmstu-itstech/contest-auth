package user

import (
	"github.com/bmstu-itstech/contest-auth/internal/repository"
	"github.com/bmstu-itstech/contest-auth/internal/service"
	"github.com/bmstu-itstech/contest-auth/pkg/db"
)

type userService struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

// NewService creates a new instance of userService with the provided UserRepository.
func NewService(
	userRepository repository.UserRepository,
	txManager db.TxManager,
) service.UserService {
	return &userService{
		userRepository: userRepository,
		txManager:      txManager,
	}
}
