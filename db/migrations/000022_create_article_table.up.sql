CREATE TABLE article (
  id serial PRIMARY KEY,
  title VARCHAR (150) NOT NULL,
  content TEXT,
  author_id INT,
  updatedAt TIMESTAMP NULL,
  createdAt TIMESTAMP NULL
);
