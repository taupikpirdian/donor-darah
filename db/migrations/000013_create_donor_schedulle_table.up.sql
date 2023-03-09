CREATE TABLE donor_schedulle (
  id serial PRIMARY KEY,
  unitId INT UNSIGNED NOT NULL,
  placeName VARCHAR (150) NOT NULL,
  address TEXT NULL,
  date DATE NOT NULL, 
  timeStart TIME NOT NULL, 
  timeEnd TIME NOT NULL, 
  type VARCHAR (50) NOT NULL, 
  updatedAt TIMESTAMP NULL,
  createdAt TIMESTAMP NULL
);