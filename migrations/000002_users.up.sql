ALTER TABLE users
DROP COLUMN first_name,
DROP COLUMN last_name,
ADD COLUMN email TEXT,
ADD COLUMN password TEXT;