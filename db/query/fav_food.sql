-- name: CreateFood :one
INSERT INTO fav_food (
  "Company", 
  "Variety",
  "Protein",
  "Fat",
  "Carbs",
  "Phos"
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetFoodByID :many
SELECT * FROM fav_food
 ORDER BY id asc
 LIMIT $1 OFFSET $2;

-- name: GetFoodByCompany :many
SELECT * FROM fav_food
WHERE "Company" = $1
 ORDER BY id asc
 LIMIT $2 OFFSET $3;

-- name: UpdateFood :one
UPDATE fav_food
SET "Company" = $2,
    "Variety" = $3,
    "Protein" = $4,
    "Fat" = $5,
    "Carbs" = $6,
    "Phos" = $7
WHERE id = $1
RETURNING *;

-- name: DeleteFood :exec
DELETE FROM fav_food 
WHERE id = $1;