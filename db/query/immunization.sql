-- name: CreateImmunization :one
INSERT INTO immunization (
  name
) VALUES (
  $1
) RETURNING *;

-- name: GetImmunizationByID :one
SELECT * FROM immunization
WHERE id = $1 LIMIT 1;

-- name: GetImmunizationByName :one
SELECT * FROM immunization
WHERE name = $1 LIMIT 1;

-- name: ListImmunizations :many
SELECT * FROM immunization
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteImmunization :exec
DELETE FROM immunization 
WHERE id = $1;