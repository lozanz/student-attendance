-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE class (
  id SERIAL PRIMARY KEY,
  name VARCHAR(256)
)

-- +migrate StatementEnd