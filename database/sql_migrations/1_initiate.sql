-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(256) NOT NULL,
  password VARCHAR(256) NOT NULL,
  role VARCHAR(256) NOT NULL
)

-- +migrate StatementEnd