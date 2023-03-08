CREATE TABLE notifications (
  id serial PRIMARY KEY,
  userId INT UNSIGNED NOT NULL,
  title VARCHAR (150) NOT NULL,
  message TEXT NOT NULL,
  updatedAt TIMESTAMP NULL,
  createdAt TIMESTAMP NULL
);