package controllers

import (
	db "happiness-to-straycat/db/sqlc"
	"net/http"

	"happiness-to-straycat/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	db *db.Queries
}

func NewAuthController(db *db.Queries) *AuthController {
	return &AuthController{db}
}

func (ac *AuthController) SignUpUser(ctx *fiber.Ctx) error {
	var credentials *db.User

	if err := ctx.BodyParser(&credentials); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	hashedPassword := utils.HashPassword(credentials.Password)

	args := &db.CreateUserParams{
		Email:    credentials.Email,
		Password: hashedPassword,
	}

	user, err := ac.db.CreateUser(ctx.Context(), *args)

	if err != nil {
		return ctx.Status(http.StatusBadGateway).JSON(err.Error())
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}
