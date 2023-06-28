CREATE TABLE blood_stock (
  id serial PRIMARY KEY,
  unitId INT UNSIGNED NOT NULL,
  title VARCHAR (150) NOT NULL,
  updatedAt TIMESTAMP NULL,
  createdAt TIMESTAMP NULL
);