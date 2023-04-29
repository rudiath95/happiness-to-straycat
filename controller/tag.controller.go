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

type TagController struct {
	db  *db.Queries
	ctx context.Context
}

func NewTagController(db *db.Queries, ctx context.Context) *TagController {
	return &TagController{db, ctx}
}

func (ac *TagController) CreateTag(c *fiber.Ctx) error {

	var payload *schemas.CreateTag

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	args := &db.CreateTagParams{
		Name:      payload.Name,
		UpdatedAt: time.Now(),
	}
	post, err := ac.db.CreateTag(c.Context(), *args)

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"tag":    post,
	})
}

func (ac *TagController) UpdateTag(c *fiber.Ctx) error {
	var payload *schemas.UpdateTag
	tagId := c.Params("tagId")
	i, _ := strconv.ParseInt(tagId, 10, 64)

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	now := time.Now()
	args := &db.UpdateTagParams{
		ID:        i,
		Name:      sql.NullString{String: payload.Name, Valid: payload.Name != ""},
		UpdatedAt: sql.NullTime{Time: now, Valid: true},
	}

	post, err := ac.db.UpdateTag(c.Context(), *args)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Tag with that ID exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "tag": post})
}

func (ac *TagController) GetTagById(c *fiber.Ctx) error {
	tagId, err := strconv.ParseInt(c.Params("tagId"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid Tag ID"})
	}

	post, err := ac.db.GetTagByID(c.Context(), tagId)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Tag with that ID exists"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "tag": post})
}

func (ac *TagController) GetAllTags(c *fiber.Ctx) error {
	page := c.Query("page", "1")
	limit := c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	args := &db.ListTagsParams{
		Limit:  int32(intLimit),
		Offset: int32(offset),
	}

	posts, err := ac.db.ListTags(c.Context(), *args)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	if posts == nil {
		posts = []db.Tag{}
	}

	return c.JSON(fiber.Map{"status": "success", "results": len(posts), "data": posts})
}

func (ac *TagController) DeleteTagById(c *fiber.Ctx) error {
	tagId := c.Params("tagId")
	i, _ := strconv.ParseInt(tagId, 10, 64)

	_, err := ac.db.GetTagByID(c.Context(), i)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  "fail",
				"message": "No Tag with that ID exists",
			})
		}
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	err = ac.db.DeleteTag(c.Context(), i)
	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.SendStatus(http.StatusNoContent)
}
