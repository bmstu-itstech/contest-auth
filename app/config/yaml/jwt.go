package yaml

import "time"

// JWT holds the configuration for the JWT.
type JWT struct {
	RefreshTokenSecretKey  string
	AccessTokenSecretKey   string
	RefreshTokenExpiration time.Duration `validate:"required" yaml:"refresh_token_expiration"`
	AccessTokenExpiration  time.Duration `validate:"required" yaml:"access_token_expiration"`
}
