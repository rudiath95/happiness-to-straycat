package models

import (
	"fmt"
	db "happiness-to-straycat/db/sqlc"
	"time"
)

// ? SignInInput struct
type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ? UserResponse struct
type UserResponse struct {
	ID        int64     `json:"id,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FilteredResponse(user db.User) UserResponse {
	var convertedString string
	if str, ok := user.Role.(string); ok {
		convertedString = fmt.Sprintf("%v", user.Role)
	} else {
		fmt.Println("fail to convert", str)
	}
	return UserResponse{
		Email:     user.Email,
		Role:      convertedString,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
