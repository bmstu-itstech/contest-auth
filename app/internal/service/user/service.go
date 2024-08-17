package user

import (
	"context"

	"github.com/bmstu-itstech/contest-auth/internal/model"
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

// Registration handles the business logic for user registration.
func (s *userService) Registration(
	ctx context.Context, params model.RegistrationParams,
) (resp model.RegistrationResponse, err error) {
	// TODO: Implement registration logic here
	return model.RegistrationResponse{}, nil
}

// Login handles the business logic for user login.
func (s *userService) Login(
	ctx context.Context, params model.LoginParams,
) (resp model.LoginResponse, err error) {
	// TODO: Implement login logic here
	return model.LoginResponse{}, nil
}

// Logout handles the business logic for user logout.
func (s *userService) Logout(
	ctx context.Context, params model.LogoutParams,
) (resp model.LogoutResponse, err error) {
	// TODO: Implement logout logic here
	return model.LogoutResponse{}, nil
}

// GetRefreshToken handles the logic to issue a new refresh token.
func (s *userService) GetRefreshToken(
	ctx context.Context, params model.GetRefreshTokenParams,
) (resp model.GetRefreshTokenResponse, err error) {
	// TODO: Implement refresh token issuance logic here
	return model.GetRefreshTokenResponse{}, nil
}

// GetAccessToken handles the logic to issue a new access token using a refresh token.
func (s *userService) GetAccessToken(
	ctx context.Context, params model.GetAccessTokenParams,
) (resp model.GetAccessTokenResponse, err error) {
	// TODO: Implement access token issuance logic here
	return model.GetAccessTokenResponse{}, nil
}
