CREATE TABLE category (
  id serial PRIMARY KEY,
  name VARCHAR (150) NOT NULL,
  tag VARCHAR (150) NOT NULL,
  updatedAt TIMESTAMP NULL,
  createdAt TIMESTAMP NULL
);
