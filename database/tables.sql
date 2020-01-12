CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    isbn text NOT NULL,
    title text NOT NULL,
    author_id integer REFERENCES author(id) ON UPDATE CASCADE
);

CREATE TABLE author (
    id SERIAL PRIMARY KEY,
    firstname text,
    lastname text
);