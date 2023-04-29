package controllers

import (
	"context"
	"database/sql"
	"errors"
	db "happiness-to-straycat/db/sqlc"
	"happiness-to-straycat/schemas"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type FavFoodController struct {
	db  *db.Queries
	ctx context.Context
}

func NewFavFoodController(db *db.Queries, ctx context.Context) *FavFoodController {
	return &FavFoodController{db, ctx}
}

func (ac *FavFoodController) CreateFavFood(c *fiber.Ctx) error {

	var payload *schemas.CreateFavFood

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	args := &db.CreateFoodParams{
		Company:   payload.Company,
		Variety:   sql.NullString{String: payload.Variety, Valid: true},
		Protein:   int32(payload.Protein),
		Fat:       int32(payload.Fat),
		Carbs:     int32(payload.Carbs),
		Phos:      int32(payload.Phos),
		Notes:     sql.NullString{String: payload.Notes, Valid: true},
		UpdatedAt: time.Now(),
	}
	post, err := ac.db.CreateFood(c.Context(), *args)

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":   "success",
		"fav_food": post,
	})
}

func (ac *FavFoodController) UpdateFavFood(c *fiber.Ctx) error {
	var payload *schemas.UpdateFavFood
	foodId := c.Params("foodId")
	i, _ := strconv.ParseInt(foodId, 10, 64)

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	now := time.Now()
	args := &db.UpdateFoodParams{
		ID:        i,
		Company:   payload.Company,
		Variety:   sql.NullString{String: payload.Variety, Valid: payload.Notes != ""},
		Protein:   int32(payload.Protein),
		Fat:       int32(payload.Fat),
		Carbs:     int32(payload.Carbs),
		Phos:      int32(payload.Phos),
		Notes:     sql.NullString{String: payload.Notes, Valid: payload.Notes != ""},
		UpdatedAt: now,
	}

	post, err := ac.db.UpdateFood(c.Context(), *args)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Food with that ID exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "food": post})
}

func (ac *FavFoodController) GetFavFoodById(c *fiber.Ctx) error {
	foodId, err := strconv.ParseInt(c.Params("foodId"), 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid Food ID"})
	}

	post, err := ac.db.GetFoodByID(c.Context(), foodId)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Food with that ID exists"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "food": post})
}

func (ac *FavFoodController) GetAllFavFoods(c *fiber.Ctx) error {
	page := c.Query("page", "1")
	limit := c.Query("limit", "20")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	args := &db.ListFavFoodsParams{
		Limit:  int32(intLimit),
		Offset: int32(offset),
	}

	posts, err := ac.db.ListFavFoods(c.Context(), *args)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	if posts == nil {
		posts = []db.FavFood{}
	}

	return c.JSON(fiber.Map{"status": "success", "results": len(posts), "data": posts})
}

func (ac *FavFoodController) DeleteFavFoodById(c *fiber.Ctx) error {
	foodId := c.Params("foodId")
	i, _ := strconv.ParseInt(foodId, 10, 64)

	_, err := ac.db.GetFoodByID(c.Context(), i)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  "fail",
				"message": "No Food with that ID exists",
			})
		}
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	err = ac.db.DeleteFood(c.Context(), i)
	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.SendStatus(http.StatusNoContent)
}
