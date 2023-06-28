CREATE TABLE blood_stock_detail (
  id serial PRIMARY KEY,
  blood_stock_id INT UNSIGNED NOT NULL,
  title VARCHAR (150) NOT NULL,
  stock INT,
  updatedAt TIMESTAMP NULL,
  createdAt TIMESTAMP NULL
);