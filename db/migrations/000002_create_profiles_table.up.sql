CREATE TABLE IF NOT EXISTS profiles(
   id serial PRIMARY KEY,
   user_id INT NOT NULL,
   job_id INT NOT NULL,
   unit_id INT NOT NULL,
   updated_at TIMESTAMP NULL,
   created_at TIMESTAMP NULL
);