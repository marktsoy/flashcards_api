CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    email VARCHAR(50) UNIQUE NOT NULL,
    encrypted_password VARCHAR(255) NOT NULL 
)