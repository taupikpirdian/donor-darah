CREATE TABLE IF NOT EXISTS regencies(
   id serial PRIMARY KEY,
   province_id INT NOT NULL,
   name VARCHAR (150) NULL,
   updatedAt TIMESTAMP NULL,
   createdAt TIMESTAMP NULL
);