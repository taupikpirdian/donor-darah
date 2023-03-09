CREATE TABLE donor_registers (
  id serial PRIMARY KEY,
  code VARCHAR (150) NOT NULL,
  userId INT UNSIGNED NOT NULL,
  donorSchedulleId INT UNSIGNED NOT NULL,
  isApprove INT DEFAULT (1),
  donorProff VARCHAR (150) NULL, 
  updatedAt TIMESTAMP NULL,
  createdAt TIMESTAMP NULL
);