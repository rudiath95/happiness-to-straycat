package models

import (
	"fmt"
	db "happiness-to-straycat/db/sqlc"
	"time"

	"github.com/google/uuid"
)

// ? SignInInput struct
type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ? UserResponse struct
type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type UserResponseID struct {
	ID uuid.UUID `json:"id,omitempty"`
}

func FilteredResponse(user *db.User) UserResponse {
	var convertedString string
	if str, ok := user.Role.(string); ok {
		convertedString = fmt.Sprintf("%v", user.Role)
	} else {
		_ = str
	}
	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Role:      convertedString,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func FilterUserID(user db.User) UserResponse {
	return UserResponse{
		ID: user.ID,
	}
}
