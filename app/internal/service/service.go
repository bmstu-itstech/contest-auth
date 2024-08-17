package service

import (
	"context"

	"github.com/bmstu-itstech/contest-auth/internal/model"
)

type UserService interface {
	// Registration handles the business logic for user registration.
	Registration(ctx context.Context, params model.RegistrationParams) (resp model.RegistrationResponse, err error)

	// Login handles the business logic for user login.
	Login(ctx context.Context, params model.LoginParams) (resp model.LoginResponse, err error)

	// Logout handles the business logic for user logout.
	Logout(ctx context.Context, params model.LogoutParams) (resp model.LogoutResponse, err error)

	// GetRefreshToken handles the logic to issue a new refresh token.
	GetRefreshToken(ctx context.Context, params model.GetRefreshTokenParams) (resp model.GetRefreshTokenResponse, err error)

	// GetAccessToken handles the logic to issue a new access token using a refresh token.
	GetAccessToken(ctx context.Context, params model.GetAccessTokenParams) (resp model.GetAccessTokenResponse, err error)
}
