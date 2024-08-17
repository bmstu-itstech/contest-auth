package user

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/bmstu-itstech/contest-auth/config"
	"github.com/bmstu-itstech/contest-auth/internal/model"
	"github.com/bmstu-itstech/contest-auth/internal/repository"
	"github.com/bmstu-itstech/contest-auth/internal/service"
	"github.com/bmstu-itstech/contest-auth/internal/utils"
	"github.com/bmstu-itstech/contest-auth/pkg/db"
)

type userService struct {
	cfg            *config.Config
	userRepository repository.UserRepository
	txManager      db.TxManager
}

// NewService creates a new instance of userService with the provided UserRepository.
func NewService(
	cfg *config.Config,
	userRepository repository.UserRepository,
	txManager db.TxManager,
) service.UserService {
	return &userService{
		cfg:            cfg,
		userRepository: userRepository,
		txManager:      txManager,
	}
}

// Registration handles the business logic for user registration.
func (s *userService) Registration(
	ctx context.Context,
	params model.RegistrationParams,
) (resp model.RegistrationResponse, err error) {
	// Check if user already exists
	existingUser, err := s.userRepository.GetUserByEmail(ctx, model.GetUserByEmailParams{
		Email: params.Email,
	})
	if err == nil && existingUser.UserID != "" {
		// User already exists
		err = errors.New("user already exists")
		return
	}

	hashedPassword, err := utils.HashPassword(params.Password)
	if err != nil {
		return
	}

	_, err = s.userRepository.CreateUser(ctx, model.CreateUserParams{
		Username:     params.Username,
		Email:        params.Email,
		PasswordHash: string(hashedPassword),
	})
	if err != nil {
		return
	}

	return model.RegistrationResponse{
		Success: true,
	}, nil
}

// Login handles the business logic for user login.
func (s *userService) Login(
	ctx context.Context,
	params model.LoginParams,
) (resp model.LoginResponse, err error) {
	user, err := s.userRepository.GetUserByEmail(ctx, model.GetUserByEmailParams{
		Email: params.Email,
	})
	if err != nil {
		return
	}

	isPasswordsEqual := utils.VerifyPassword(user.PasswordHash, params.Password)
	if !isPasswordsEqual {
		err = errors.New("invalid password")
		return
	}

	refreshToken, err := utils.GenerateToken(model.UserInfoJWT{
		Email: params.Email,
	},
		[]byte(s.cfg.JWT.RefreshTokenSecretKey),
		s.cfg.JWT.RefreshTokenExpiration,
	)
	if err != nil {
		err = errors.New("failed to generate token")
		return
	}

	return model.LoginResponse{
		RefreshToken: refreshToken,
	}, nil
}

// Logout handles the business logic for user logout.
func (s *userService) Logout(
	ctx context.Context,
	params model.LogoutParams,
) (resp model.LogoutResponse, err error) {
	_, err = utils.VerifyToken(params.RefreshToken, []byte(s.cfg.JWT.RefreshTokenSecretKey))
	if err != nil {
		err = status.Errorf(codes.Aborted, "invalid refresh token")
		return
	}

	return model.LogoutResponse{
		RefreshToken: "",
	}, nil
}

// GetRefreshToken handles the logic to issue a new refresh token.
func (s *userService) GetRefreshToken(
	ctx context.Context,
	params model.GetRefreshTokenParams,
) (resp model.GetRefreshTokenResponse, err error) {
	claims, err := utils.VerifyToken(params.RefreshToken, []byte(s.cfg.JWT.RefreshTokenSecretKey))
	if err != nil {
		err = status.Errorf(codes.Aborted, "invalid refresh token")
		return
	}

	refreshToken, err := utils.GenerateToken(model.UserInfoJWT{
		Email: claims.Email,
	},
		[]byte(s.cfg.JWT.RefreshTokenSecretKey),
		s.cfg.JWT.RefreshTokenExpiration,
	)
	if err != nil {
		return
	}

	return model.GetRefreshTokenResponse{
		RefreshToken: refreshToken,
	}, nil
}

// GetAccessToken handles the logic to issue a new access token using a refresh token.
func (s *userService) GetAccessToken(
	ctx context.Context,
	params model.GetAccessTokenParams,
) (resp model.GetAccessTokenResponse, err error) {
	claims, err := utils.VerifyToken(params.RefreshToken, []byte(s.cfg.JWT.RefreshTokenSecretKey))
	if err != nil {
		err = status.Errorf(codes.Aborted, "invalid refresh token")
		return
	}

	accessToken, err := utils.GenerateToken(model.UserInfoJWT{
		Email: claims.Email,
	},
		[]byte(s.cfg.JWT.AccessTokenSecretKey),
		s.cfg.JWT.AccessTokenExpiration,
	)
	if err != nil {
		return
	}

	return model.GetAccessTokenResponse{
		AccessToken: accessToken,
	}, nil
}
