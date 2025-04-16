-- name: AddUser :one
INSERT INTO users (username , email , password_hashed)
VALUES ($1 , $2 , $3)
RETURNING id, username, email , password_hashed , created_at;

-- name: GetUserByUsername :one
SELECT id, username, email, password_hashed, created_at
FROM users 
WHERE username = $1;
