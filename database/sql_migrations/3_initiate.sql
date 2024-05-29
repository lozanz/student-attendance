-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE student (
  id SERIAL PRIMARY KEY ,
  name VARCHAR(256) NOT NULL,
  jenis_kelamin VARCHAR(10) NOT NULL,
  alamat Text NOT NULL,
  class_id INT REFERENCES class(id)
)

-- +migrate StatementEnd