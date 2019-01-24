
-- +migrate Up

CREATE TABLE categories (
  id serial NOT NULL primary key,
  category  text NOT NULL,
  timestamp timestamp without time zone DEFAULT now() NOT NULL
);

-- +migrate Down

DROP TABLE IF EXISTS categories;
