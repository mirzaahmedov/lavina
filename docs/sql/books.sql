CREATE TABLE books (
  id serial PRIMARY KEY,
  title VARCHAR ( 250 ) NOT NULL,
  author VARCHAR ( 250 ) NOT NULL,
  published VARCHAR ( 250 ) NOT NULL,
  pages numeric NOT NULL,
  status smallint NOT NULL,
  isbn VARCHAR ( 20 ) NOT NULL UNIQUE,
  user_id numeric NOT NULL
)

SELECT id, title, author, published, pages, status, isbn FROM books WHERE user_id=$1;
DELETE FROM books WHERE id=$1 AND user_id=$2