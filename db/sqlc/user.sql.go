// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: user.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  email, 
  verified,
  password,
  role,
  updated_at
) VALUES (
  $1, $2, $3, $4,$5
) RETURNING id, email, verified, password, role, created_at, updated_at
`

type CreateUserParams struct {
	Email     string      `json:"email"`
	Verified  bool        `json:"verified"`
	Password  string      `json:"password"`
	Role      interface{} `json:"role"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Email,
		arg.Verified,
		arg.Password,
		arg.Role,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Verified,
		&i.Password,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUserAndDetail = `-- name: CreateUserAndDetail :one
WITH inserted_user AS (
  INSERT INTO users (
    email, 
    verified,
    password,
    role,
    updated_at
  ) VALUES ($1, $2, $3, $4, $5)
  RETURNING id, email, verified, password, role, created_at, updated_at
)
INSERT INTO user_detail (
  user_id, 
  name, 
  gender, 
  age, 
  address, 
  phone, 
  created_at, 
  updated_at
) 
VALUES (
  (SELECT id FROM inserted_user), 
  $6, 
  $7, 
  $8, 
  $9, 
  $10, 
  $11, 
  $12
)
RETURNING id, user_id, name, gender, age, address, phone, created_at, updated_at
`

type CreateUserAndDetailParams struct {
	Email       string         `json:"email"`
	Verified    bool           `json:"verified"`
	Password    string         `json:"password"`
	Role        interface{}    `json:"role"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Name        sql.NullString `json:"name"`
	Gender      interface{}    `json:"gender"`
	Age         sql.NullInt32  `json:"age"`
	Address     sql.NullString `json:"address"`
	Phone       sql.NullInt32  `json:"phone"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt_2 time.Time      `json:"updated_at_2"`
}

func (q *Queries) CreateUserAndDetail(ctx context.Context, arg CreateUserAndDetailParams) (UserDetail, error) {
	row := q.db.QueryRowContext(ctx, createUserAndDetail,
		arg.Email,
		arg.Verified,
		arg.Password,
		arg.Role,
		arg.UpdatedAt,
		arg.Name,
		arg.Gender,
		arg.Age,
		arg.Address,
		arg.Phone,
		arg.CreatedAt,
		arg.UpdatedAt_2,
	)
	var i UserDetail
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Gender,
		&i.Age,
		&i.Address,
		&i.Phone,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUserDetail = `-- name: CreateUserDetail :one
INSERT INTO user_detail (
  user_id, 
  name,
  gender,
  age,
  address,
  phone
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id, user_id, name, gender, age, address, phone, created_at, updated_at
`

type CreateUserDetailParams struct {
	UserID  uuid.UUID      `json:"user_id"`
	Name    sql.NullString `json:"name"`
	Gender  interface{}    `json:"gender"`
	Age     sql.NullInt32  `json:"age"`
	Address sql.NullString `json:"address"`
	Phone   sql.NullInt32  `json:"phone"`
}

func (q *Queries) CreateUserDetail(ctx context.Context, arg CreateUserDetailParams) (UserDetail, error) {
	row := q.db.QueryRowContext(ctx, createUserDetail,
		arg.UserID,
		arg.Name,
		arg.Gender,
		arg.Age,
		arg.Address,
		arg.Phone,
	)
	var i UserDetail
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Gender,
		&i.Age,
		&i.Address,
		&i.Phone,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, email, verified, password, role, created_at, updated_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Verified,
		&i.Password,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, verified, password, role, created_at, updated_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Verified,
		&i.Password,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserDetail = `-- name: GetUserDetail :one
SELECT id, user_id, name, gender, age, address, phone, created_at, updated_at FROM user_detail
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetUserDetail(ctx context.Context, userID uuid.UUID) (UserDetail, error) {
	row := q.db.QueryRowContext(ctx, getUserDetail, userID)
	var i UserDetail
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Gender,
		&i.Age,
		&i.Address,
		&i.Phone,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, email, verified, password, role, created_at, updated_at FROM users
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Verified,
			&i.Password,
			&i.Role,
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

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET password = $2
WHERE id = $1
RETURNING id, email, verified, password, role, created_at, updated_at
`

type UpdateUserParams struct {
	ID       uuid.UUID `json:"id"`
	Password string    `json:"password"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.ID, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Verified,
		&i.Password,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserDetail = `-- name: UpdateUserDetail :one
UPDATE user_detail
SET name = $2,
name = $3,
gender = $4,
age = $5,
address = $6,
phone = $7
WHERE user_id = $1
RETURNING id, user_id, name, gender, age, address, phone, created_at, updated_at
`

type UpdateUserDetailParams struct {
	UserID  uuid.UUID      `json:"user_id"`
	Name    sql.NullString `json:"name"`
	Name_2  sql.NullString `json:"name_2"`
	Gender  interface{}    `json:"gender"`
	Age     sql.NullInt32  `json:"age"`
	Address sql.NullString `json:"address"`
	Phone   sql.NullInt32  `json:"phone"`
}

func (q *Queries) UpdateUserDetail(ctx context.Context, arg UpdateUserDetailParams) (UserDetail, error) {
	row := q.db.QueryRowContext(ctx, updateUserDetail,
		arg.UserID,
		arg.Name,
		arg.Name_2,
		arg.Gender,
		arg.Age,
		arg.Address,
		arg.Phone,
	)
	var i UserDetail
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Gender,
		&i.Age,
		&i.Address,
		&i.Phone,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
