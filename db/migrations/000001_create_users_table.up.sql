CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   name VARCHAR (150) NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL,
   phone VARCHAR (20) UNIQUE NOT NULL,
   password VARCHAR (50) NOT NULL,
   updated_at TIMESTAMP NULL,
   created_at TIMESTAMP NULL
);