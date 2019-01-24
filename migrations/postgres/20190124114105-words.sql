
-- +migrate Up

CREATE TABLE words (
  id serial NOT NULL primary key,
  word text NOT NULL,
  author text NOT NULL,
  category_id integer NOT NULL references categories(id),
  timestamp timestamp without time zone DEFAULT now() NOT NULL
);

-- +migrate Down

DROP TABLE IF EXISTS words;
