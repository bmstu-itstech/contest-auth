-- +goose Up
CREATE TABLE users (
  id uuid DEFAULT gen_random_uuid(),
  username varchar(255) NOT NULL,
  email varchar(255) UNIQUE NOT NULL,
  password_hash varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),

  PRIMARY KEY (id)
);


-- +goose Down
DROP TABLE users;
