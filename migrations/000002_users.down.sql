ALTER TABLE users
ADD COLUMN first_name TEXT,
ADD COLUMN last_name TEXT,
DROP COLUMN email,
DROP COLUMN password;