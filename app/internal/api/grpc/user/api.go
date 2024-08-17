package user

import (
	"context"

	"github.com/bmstu-itstech/contest-auth/internal/converter"
	"github.com/bmstu-itstech/contest-auth/internal/service"
	pb "github.com/bmstu-itstech/contest-auth/pkg/user_v1"
)

// GRPCHandlers represents the gRPC handlers that implement the UserV1Server interface
// and use the UserService for business logic operations.
type GRPCHandlers struct {
	pb.UnimplementedUserV1Server
	userService service.UserService
}

// NewGRPCHandlers creates a new instance of GRPCHandlers with the provided UserService.
func NewGRPCHandlers(userService service.UserService) *GRPCHandlers {
	return &GRPCHandlers{
		userService: userService,
	}
}

// Registration handles user registration requests.
func (h *GRPCHandlers) Registration(ctx context.Context, req *pb.RegistrationRequest) (*pb.RegistrationResponse, error) {
	resp, err := h.userService.Registration(ctx, converter.ConvertRegistrationParamsFromHandlerToService(req))
	if err != nil {
		return nil, err
	}

	return converter.ConvertRegistrationResponseFromServiceToHandler(resp), nil
}

// Login handles user login requests.
func (h *GRPCHandlers) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	resp, err := h.userService.Login(ctx, converter.ConvertLoginParamsFromHandlerToService(req))
	if err != nil {
		return nil, err
	}

	return converter.ConvertLoginResponseFromServiceToHandler(resp), nil
}

// Logout handles user logout requests.
func (h *GRPCHandlers) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	resp, err := h.userService.Logout(ctx, converter.ConvertLogoutParamsFromHandlerToService(req))
	if err != nil {
		return nil, err
	}

	return converter.ConvertLogoutResponseFromServiceToHandler(resp), nil
}

// GetRefreshToken handles requests to obtain a new refresh token.
func (h *GRPCHandlers) GetRefreshToken(ctx context.Context, req *pb.GetRefreshTokenRequest) (*pb.GetRefreshTokenResponse, error) {
	resp, err := h.userService.GetRefreshToken(ctx, converter.ConvertGetRefreshTokenParamsFromHandlerToService(req))
	if err != nil {
		return nil, err
	}

	return converter.ConvertGetRefreshTokenResponseFromServiceToHandler(resp), nil
}

// GetAccessToken handles requests to obtain a new access token using a refresh token.
func (h *GRPCHandlers) GetAccessToken(ctx context.Context, req *pb.GetAccessTokenRequest) (*pb.GetAccessTokenResponse, error) {
	resp, err := h.userService.GetAccessToken(ctx, converter.ConvertGetAccessTokenParamsFromHandlerToService(req))
	if err != nil {
		return nil, err
	}

	return converter.ConvertGetAccessTokenResponseFromServiceToHandler(resp), nil
}
