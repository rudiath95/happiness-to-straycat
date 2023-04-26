package controllers

import (
	"context"
	db "happiness-to-straycat/db/sqlc"
	"happiness-to-straycat/models"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	db  *db.Queries
	ctx context.Context
}

func NewUserController(db *db.Queries, ctx context.Context) UserController {
	return UserController{db, ctx}
}

func (uc *UserController) GetMe(ctx *fiber.Ctx) error {
	currentUser := ctx.Locals("currentUser").(db.User)

	return ctx.JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": models.FilteredResponse(currentUser)}})
}
