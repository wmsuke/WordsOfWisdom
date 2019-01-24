
-- +migrate Up

CREATE TABLE nices (
  id serial NOT NULL primary key,
  word_id integer NOT NULL references words(id),
  user_id integer NOT NULL references users(id)
);

-- +migrate Down

DROP TABLE IF EXISTS nices;
