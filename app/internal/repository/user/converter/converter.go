package converter

import (
	"github.com/bmstu-itstech/contest-auth/internal/model"
	modelRepo "github.com/bmstu-itstech/contest-auth/internal/repository/user/model"
)

// ConvertCreateUserParamsFromServiceToRepo converts CreateUserParams from the service layer to the repository layer.
func ConvertCreateUserParamsFromServiceToRepo(params model.CreateUserParams) modelRepo.CreateUserParams {
	return modelRepo.CreateUserParams(params)
}

// ConvertCreateUserResponseFromRepoToService converts CreateUserResponse from the repository layer to the service layer.
func ConvertCreateUserResponseFromRepoToService(params modelRepo.CreateUserResponse) model.CreateUserResponse {
	return model.CreateUserResponse(params)
}

// ConvertGetUserByEmailParamsFromServiceToRepo converts GetUserByEmailParams from the service layer to the repository layer.
func ConvertGetUserByEmailParamsFromServiceToRepo(params model.GetUserByEmailParams) modelRepo.GetUserByEmailParams {
	return modelRepo.GetUserByEmailParams(params)
}

// ConvertGetUserByEmailResponseFromRepoToService converts GetUserByEmailResponse from the repository layer to the service layer.
func ConvertGetUserByEmailResponseFromRepoToService(params modelRepo.GetUserByEmailResponse) model.GetUserByEmailResponse {
	return model.GetUserByEmailResponse(params)
}
