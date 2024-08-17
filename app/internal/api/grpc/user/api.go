package user

import (
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
