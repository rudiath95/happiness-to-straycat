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

type ImmunizationController struct {
	db  *db.Queries
	ctx context.Context
}

func NewImmunizationController(db *db.Queries, ctx context.Context) *ImmunizationController {
	return &ImmunizationController{db, ctx}
}

func (ac *ImmunizationController) CreateImmunization(c *fiber.Ctx) error {

	var payload *schemas.CreateImmunization

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	name := sql.NullString{String: payload.Name, Valid: true}

	args := &db.CreateImmunizationParams{
		Name:      name,
		UpdatedAt: time.Now(),
	}
	post, err := ac.db.CreateImmunization(c.Context(), *args)

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":       "success",
		"immunization": post,
	})
}

func (ac *ImmunizationController) UpdateImmunization(c *fiber.Ctx) error {
	var payload *schemas.UpdateImmunization
	immuneId := c.Params("immuneId")
	i, _ := strconv.ParseInt(immuneId, 10, 64)

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	now := time.Now()
	args := &db.UpdateImmunizationParams{
		ID:        i,
		Name:      sql.NullString{String: payload.Name, Valid: payload.Name != ""},
		UpdatedAt: sql.NullTime{Time: now, Valid: true},
	}

	post, err := ac.db.UpdateImmunization(c.Context(), *args)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Immunization with that ID exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "immunization": post})
}

func (ac *ImmunizationController) GetImmunizationById(c *fiber.Ctx) error {
	immuneID, err := strconv.ParseInt(c.Params("immuneId"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid Immunization ID"})
	}

	post, err := ac.db.GetImmunizationByID(c.Context(), immuneID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Immunization with that ID exists"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "immunization": post})
}

func (ac *ImmunizationController) GetAllImmunizations(c *fiber.Ctx) error {
	page := c.Query("page", "1")
	limit := c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	args := &db.ListImmunizationsParams{
		Limit:  int32(intLimit),
		Offset: int32(offset),
	}

	posts, err := ac.db.ListImmunizations(c.Context(), *args)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	if posts == nil {
		posts = []db.Immunization{}
	}

	return c.JSON(fiber.Map{"status": "success", "results": len(posts), "data": posts})
}

func (ac *ImmunizationController) DeleteImmunizationById(c *fiber.Ctx) error {
	immuneId := c.Params("immuneId")
	i, _ := strconv.ParseInt(immuneId, 10, 64)

	_, err := ac.db.GetImmunizationByID(c.Context(), i)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  "fail",
				"message": "No Immunization with that ID exists",
			})
		}
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	err = ac.db.DeleteImmunization(c.Context(), i)
	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.SendStatus(http.StatusNoContent)
}
