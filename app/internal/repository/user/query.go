package user

const (
	queryCreateUser = `
		INSERT INTO	users
			(email, username, password_hash)
		VALUES
			($1, $2, $3)
		RETURNING id;
	`

	queryGetUserByEmail = `
		SELECT 
			id
			, email
			, username
			, password_hash
		FROM users
		WHERE email = $1;
	`
)
