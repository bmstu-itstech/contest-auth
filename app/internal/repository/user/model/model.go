package model

// CreateUserParams contains parameters for the repo method CreateUser method.
type CreateUserParams struct {
	Email        string `db:"email"`
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
}

// CreateUserResponse contains the response for the repo method CreateUser method.
type CreateUserResponse struct {
	UserID string `db:"id"`
}

// GetUserByEmailParams contains parameters for the repo method GetUserByEmail method.
type GetUserByEmailParams struct {
	Email string `db:"email"`
}

// GetUserByEmailResponse contains the response for the repo method GetUserByEmail method.
type GetUserByEmailResponse struct {
	UserID       string `db:"id"`
	Email        string `db:"email"`
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
}
