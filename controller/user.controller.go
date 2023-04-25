package controller

import (
	"context"
	db "happiness-to-straycat/db/sqlc"
	"happiness-to-straycat/schemas"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	db  *db.Queries
	ctx context.Context
}

func NewUserController(db *db.Queries, ctx context.Context) *UserController {
	return &UserController{db, ctx}
}

func (uc *UserController) CreateUser(ctx *fiber.Ctx) error {
	var payload *schemas.CreateUser

	if err := ctx.BodyParser(&payload); err != nil {
		ctx.Status(503).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
		return err
	}

	args := &db.CreateUserParams{
		Email:    payload.Email,
		Password: payload.Password,
	}

	user, err := uc.db.CreateUser(ctx.Context(), *args)

	if err != nil {
		ctx.Status(503).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
		return err
	}

	ctx.JSON(fiber.Map{
		"message": "Create Data Successfully",
		"user":    user,
	})

	return nil
}
