package model

// RegistrationParams contains parameters for the Registration method.
type RegistrationParams struct {
	Email    string
	Username string
	Password string
}

// RegistrationResponse contains the response for the Registration method.
type RegistrationResponse struct {
	Success bool
}

// LoginParams contains parameters for the Login method.
type LoginParams struct {
	Email    string
	Password string
}

// LoginResponse contains the response for the Login method.
type LoginResponse struct {
	RefreshToken string
}

// LogoutParams contains parameters for the Logout method.
type LogoutParams struct {
	RefreshToken string
}

// LogoutResponse contains the response for the Logout method.
type LogoutResponse struct {
	RefreshToken string
}

// GetRefreshTokenParams contains parameters for the GetRefreshToken method.
type GetRefreshTokenParams struct {
	RefreshToken string
}

// GetRefreshTokenResponse contains the response for the GetRefreshToken method.
type GetRefreshTokenResponse struct {
	RefreshToken string
}

// GetAccessTokenParams contains parameters for the GetAccessToken method.
type GetAccessTokenParams struct {
	RefreshToken string
}

// GetAccessTokenResponse contains the response for the GetAccessToken method.
type GetAccessTokenResponse struct {
	AccessToken string
}

// CreateUserParams contains parameters for the CreateUser method.
type CreateUserParams struct {
	Email        string
	Username     string
	PasswordHash string
}

// CreateUserResponse contains the response for the CreateUser method.
type CreateUserResponse struct {
	UserID string
}

// GetUserByEmailParams contains parameters for the GetUserByEmail method.
type GetUserByEmailParams struct {
	Email string
}

// GetUserByEmailResponse contains the response for the GetUserByEmail method.
type GetUserByEmailResponse struct {
	UserID       string
	Email        string
	Username     string
	PasswordHash string
}
