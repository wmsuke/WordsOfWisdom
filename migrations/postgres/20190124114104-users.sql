
-- +migrate Up

CREATE TABLE users (
  id serial NOT NULL primary key,
  key text NOT NULL,
  access_date timestamp NOT NULL
);

-- +migrate Down

DROP TABLE IF EXISTS users;
