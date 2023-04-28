// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: fav_food.sql

package db

import (
	"context"
	"database/sql"
)

const createFood = `-- name: CreateFood :one
INSERT INTO fav_food (
  "Company", 
  "Variety",
  "Protein",
  "Fat",
  "Carbs",
  "Phos"
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id, "Company", "Variety", "Protein", "Fat", "Carbs", "Phos", "Notes", created_at, updated_at
`

type CreateFoodParams struct {
	Company string         `json:"Company"`
	Variety sql.NullString `json:"Variety"`
	Protein int32          `json:"Protein"`
	Fat     int32          `json:"Fat"`
	Carbs   int32          `json:"Carbs"`
	Phos    int32          `json:"Phos"`
}

func (q *Queries) CreateFood(ctx context.Context, arg CreateFoodParams) (FavFood, error) {
	row := q.db.QueryRowContext(ctx, createFood,
		arg.Company,
		arg.Variety,
		arg.Protein,
		arg.Fat,
		arg.Carbs,
		arg.Phos,
	)
	var i FavFood
	err := row.Scan(
		&i.ID,
		&i.Company,
		&i.Variety,
		&i.Protein,
		&i.Fat,
		&i.Carbs,
		&i.Phos,
		&i.Notes,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteFood = `-- name: DeleteFood :exec
DELETE FROM fav_food 
WHERE id = $1
`

func (q *Queries) DeleteFood(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteFood, id)
	return err
}

const getFoodByCompany = `-- name: GetFoodByCompany :many
SELECT id, "Company", "Variety", "Protein", "Fat", "Carbs", "Phos", "Notes", created_at, updated_at FROM fav_food
WHERE "Company" = $1
 ORDER BY id asc
 LIMIT $2 OFFSET $3
`

type GetFoodByCompanyParams struct {
	Company string `json:"Company"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

func (q *Queries) GetFoodByCompany(ctx context.Context, arg GetFoodByCompanyParams) ([]FavFood, error) {
	rows, err := q.db.QueryContext(ctx, getFoodByCompany, arg.Company, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FavFood
	for rows.Next() {
		var i FavFood
		if err := rows.Scan(
			&i.ID,
			&i.Company,
			&i.Variety,
			&i.Protein,
			&i.Fat,
			&i.Carbs,
			&i.Phos,
			&i.Notes,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFoodByID = `-- name: GetFoodByID :many
SELECT id, "Company", "Variety", "Protein", "Fat", "Carbs", "Phos", "Notes", created_at, updated_at FROM fav_food
 ORDER BY id asc
 LIMIT $1 OFFSET $2
`

type GetFoodByIDParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetFoodByID(ctx context.Context, arg GetFoodByIDParams) ([]FavFood, error) {
	rows, err := q.db.QueryContext(ctx, getFoodByID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FavFood
	for rows.Next() {
		var i FavFood
		if err := rows.Scan(
			&i.ID,
			&i.Company,
			&i.Variety,
			&i.Protein,
			&i.Fat,
			&i.Carbs,
			&i.Phos,
			&i.Notes,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateFood = `-- name: UpdateFood :one
UPDATE fav_food
SET "Company" = $2,
    "Variety" = $3,
    "Protein" = $4,
    "Fat" = $5,
    "Carbs" = $6,
    "Phos" = $7
WHERE id = $1
RETURNING id, "Company", "Variety", "Protein", "Fat", "Carbs", "Phos", "Notes", created_at, updated_at
`

type UpdateFoodParams struct {
	ID      int64          `json:"id"`
	Company string         `json:"Company"`
	Variety sql.NullString `json:"Variety"`
	Protein int32          `json:"Protein"`
	Fat     int32          `json:"Fat"`
	Carbs   int32          `json:"Carbs"`
	Phos    int32          `json:"Phos"`
}

func (q *Queries) UpdateFood(ctx context.Context, arg UpdateFoodParams) (FavFood, error) {
	row := q.db.QueryRowContext(ctx, updateFood,
		arg.ID,
		arg.Company,
		arg.Variety,
		arg.Protein,
		arg.Fat,
		arg.Carbs,
		arg.Phos,
	)
	var i FavFood
	err := row.Scan(
		&i.ID,
		&i.Company,
		&i.Variety,
		&i.Protein,
		&i.Fat,
		&i.Carbs,
		&i.Phos,
		&i.Notes,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}