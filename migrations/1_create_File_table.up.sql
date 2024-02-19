CREATE TABLE File (
  id SERIAL PRIMARY KEY,
  filepath text NOT NULL,
  tag text NOT NULL,
  timestamp timestamp default current_timestamp
);