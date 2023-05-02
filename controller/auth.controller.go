package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"happiness-to-straycat/config"
	db "happiness-to-straycat/db/sqlc"
	"happiness-to-straycat/models"
	"happiness-to-straycat/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Role string

const (
	User  Role = "user"
	Admin Role = "admin"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type UserWithDetail struct {
	User       *db.User       `json:"user"`
	UserDetail *db.UserDetail `json:"user_detail"`
}

type AuthController struct {
	db  *db.Queries
	ctx context.Context
}

func NewAuthController(db *db.Queries, ctx context.Context) *AuthController {
	return &AuthController{db, ctx}
}
func (ac *AuthController) SignUpUser(c *fiber.Ctx) error {
	var credentials map[string]interface{}

	if err := c.BodyParser(&credentials); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	email, ok := credentials["email"].(string)
	if !ok {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid email"})
	}

	password, ok := credentials["password"].(string)
	if !ok {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid password"})
	}

	gender, ok := credentials["gender"].(string)
	if !ok {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid gender"})
	}

	age, ok := credentials["age"].(float64)
	if !ok {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid age"})
	}

	address, ok := credentials["address"].(string)
	if !ok {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid address"})
	}

	phone, ok := credentials["phone"].(float64)
	if !ok {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid phone"})
	}

	name := sql.NullString{}
	if n, ok := credentials["name"].(string); ok {
		name = sql.NullString{String: n, Valid: true}
	}

	hashedPassword := utils.HashPassword(password)

	args := db.CreateUserAndDetailParams{
		Email:     email,
		Verified:  true,
		Password:  hashedPassword,
		Role:      Role("user"),
		Name:      name,
		Gender:    Gender(gender),
		Age:       sql.NullInt32{Int32: int32(age), Valid: true},
		Address:   sql.NullString{String: address, Valid: true},
		Phone:     sql.NullInt32{Int32: int32(phone), Valid: true},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userDetail, err := ac.db.CreateUserAndDetail(ac.ctx, args)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	user := &db.User{
		ID:        userDetail.UserID,
		Email:     email,
		Verified:  true,
		Role:      Role("user"),
		UpdatedAt: time.Now(),
	}

	userResponse := models.FilteredResponse(user)

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user":        userResponse,
			"user_detail": userDetail,
		},
	})

}

// func (ac *AuthController) SignUpUser(c *fiber.Ctx) error {
// 	var credentials *db.CreateUserParams

// 	if err := c.BodyParser(&credentials); err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	hashedPassword := utils.HashPassword(credentials.Password)

// 	args := &db.CreateUserParams{
// 		Email:     credentials.Email,
// 		Password:  hashedPassword,
// 		Verified:  true,
// 		Role:      "user",
// 		UpdatedAt: time.Now(),
// 	}

// 	user, err := ac.db.CreateUser(c.Context(), *args)

// 	if err != nil {
// 		return c.Status(http.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	userResponse := models.FilteredResponse(user)

// 	return c.Status(http.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": userResponse}})
// }

func (ac *AuthController) SignInUser(c *fiber.Ctx) error {
	var credentials *models.SignInInput

	if err := c.BodyParser(&credentials); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	user, err := ac.db.GetUserByEmail(c.Context(), credentials.Email)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid email or password"})
	}

	if err := utils.ComparePassword(user.Password, credentials.Password); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid email or password"})
	}

	config, _ := config.LoadConfig(".")

	// Generate Tokens
	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    access_token,
		MaxAge:   config.AccessTokenMaxAge * 60,
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HTTPOnly: true,
		SameSite: "Lax",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refresh_token,
		MaxAge:   config.RefreshTokenMaxAge * 60,
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HTTPOnly: true,
		SameSite: "Lax",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "true",
		MaxAge:   config.AccessTokenMaxAge * 60,
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HTTPOnly: false,
		SameSite: "Lax",
	})

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": "success", "access_token": access_token})
}

func (ac *AuthController) RefreshAccessToken(ctx *fiber.Ctx) error {
	message := "could not refresh access token"

	cookie := ctx.Cookies("refresh_token")
	var err error

	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": message})
	}

	config, _ := config.LoadConfig(".")

	sub, err := utils.ValidateToken(cookie, config.RefreshTokenPublicKey)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	user, err := ac.db.GetUser(ac.ctx, uuid.MustParse(fmt.Sprint(sub)))
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "the user belonging to this token no logger exists"})
	}

	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	ctx.Cookie(&fiber.Cookie{Name: "access_token", Value: access_token, MaxAge: config.AccessTokenMaxAge * 60, Path: "/", Domain: "localhost", HTTPOnly: true, Secure: false})
	ctx.Cookie(&fiber.Cookie{Name: "logged_in", Value: "true", MaxAge: config.AccessTokenMaxAge * 60, Path: "/", Domain: "localhost", HTTPOnly: false, Secure: false})

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"status": "success", "access_token": access_token})
}

func (ac *AuthController) LogoutUser(c *fiber.Ctx) error {

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		SameSite: "Lax",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		SameSite: "Lax",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		SameSite: "Lax",
	})

	return c.JSON(fiber.Map{
		"status": "success",
	})
}
