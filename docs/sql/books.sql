CREATE TABLE books (
  id serial PRIMARY KEY,
  title VARCHAR ( 250 ) NOT NULL,
  author VARCHAR ( 250 ) NOT NULL,
  published smallint NOT NULL,
  pages numeric NOT NULL,
  status smallint NOT NULL,
  isbn VARCHAR ( 20 ) NOT NULL UNIQUE
)