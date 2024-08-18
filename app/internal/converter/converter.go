package converter

import (
	"github.com/bmstu-itstech/contest-auth/internal/model"
	pb "github.com/bmstu-itstech/contest-auth/pkg/user_v1"
)

// ConvertRegistrationParamsFromHandlerToService converts gRPC request parameters to service parameters.
func ConvertRegistrationParamsFromHandlerToService(params *pb.RegistrationRequest) model.RegistrationParams {
	return model.RegistrationParams{
		Email:    params.Email,
		Username: params.Username,
		Password: params.Password,
	}
}

// ConvertRegistrationResponseFromServiceToHandler converts service response to gRPC response.
func ConvertRegistrationResponseFromServiceToHandler(params model.RegistrationResponse) *pb.RegistrationResponse {
	return &pb.RegistrationResponse{
		Success: params.Success,
	}
}

// ConvertLoginParamsFromHandlerToService converts gRPC request parameters to service parameters.
func ConvertLoginParamsFromHandlerToService(params *pb.LoginRequest) model.LoginParams {
	return model.LoginParams{
		Email:    params.Email,
		Password: params.Password,
	}
}

// ConvertLoginResponseFromServiceToHandler converts service response to gRPC response.
func ConvertLoginResponseFromServiceToHandler(params model.LoginResponse) *pb.LoginResponse {
	return &pb.LoginResponse{
		RefreshToken: params.RefreshToken,
	}
}

// ConvertLogoutParamsFromHandlerToService converts gRPC request parameters to service parameters.
func ConvertLogoutParamsFromHandlerToService(params *pb.LogoutRequest) model.LogoutParams {
	return model.LogoutParams{
		RefreshToken: params.RefreshToken,
	}
}

// ConvertLogoutResponseFromServiceToHandler converts service response to gRPC response.
func ConvertLogoutResponseFromServiceToHandler(params model.LogoutResponse) *pb.LogoutResponse {
	return &pb.LogoutResponse{
		RefreshToken: params.RefreshToken,
	}
}

// ConvertGetRefreshTokenParamsFromHandlerToService converts gRPC request parameters to service parameters.
func ConvertGetRefreshTokenParamsFromHandlerToService(params *pb.GetRefreshTokenRequest) model.GetRefreshTokenParams {
	return model.GetRefreshTokenParams{
		RefreshToken: params.RefreshToken,
	}
}

// ConvertGetRefreshTokenResponseFromServiceToHandler converts service response to gRPC response.
func ConvertGetRefreshTokenResponseFromServiceToHandler(params model.GetRefreshTokenResponse) *pb.GetRefreshTokenResponse {
	return &pb.GetRefreshTokenResponse{
		RefreshToken: params.RefreshToken,
	}
}

// ConvertGetAccessTokenParamsFromHandlerToService converts gRPC request parameters to service parameters.
func ConvertGetAccessTokenParamsFromHandlerToService(params *pb.GetAccessTokenRequest) model.GetAccessTokenParams {
	return model.GetAccessTokenParams{
		RefreshToken: params.RefreshToken,
	}
}

// ConvertGetAccessTokenResponseFromServiceToHandler converts service response to gRPC response.
func ConvertGetAccessTokenResponseFromServiceToHandler(params model.GetAccessTokenResponse) *pb.GetAccessTokenResponse {
	return &pb.GetAccessTokenResponse{
		AccessToken: params.AccessToken,
	}
}
