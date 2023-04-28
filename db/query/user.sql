-- name: CreateUser :one
INSERT INTO users (
  email, 
  verified,
  password,
  role,
  updated_at
) VALUES (
  $1, $2, $3, $4,$5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET password = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = $1;

-- name: CreateUserDetail :one
INSERT INTO user_detail (
  user_id, 
  name,
  gender,
  age,
  address,
  phone
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetUserDetail :one
SELECT * FROM user_detail
WHERE user_id = $1 LIMIT 1;

-- name: UpdateUserDetail :one
UPDATE user_detail
SET name = $2,
name = $3,
gender = $4,
age = $5,
address = $6,
phone = $7
WHERE user_id = $1
RETURNING *;