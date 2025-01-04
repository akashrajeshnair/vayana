-- Drop the table if it already exists (useful for re-running migrations during development)
DROP TABLE IF EXISTS users;

-- Create the users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,             -- Unique identifier for each user
    username VARCHAR(255) NOT NULL UNIQUE, -- Unique username for the user
    email VARCHAR(255) NOT NULL UNIQUE,    -- Unique email for the user
    password_hash TEXT NOT NULL,           -- Hashed password for authentication
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp for when the user was created
);
