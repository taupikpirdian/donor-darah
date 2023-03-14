CREATE TABLE donor_register_questioners (
  id serial PRIMARY KEY,
  donorRegisterId INT UNSIGNED NOT NULL,
  codeQuestion VARCHAR (150) NOT NULL,
  title TEXT NOT NULL, 
  answer VARCHAR (150) NOT NULL, 
  updatedAt TIMESTAMP NULL,
  createdAt TIMESTAMP NULL
);