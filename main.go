package main

import (
	"context"
	"database/sql"
	"happiness-to-straycat/controller"
	"happiness-to-straycat/ini"
	"happiness-to-straycat/routes"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	app *fiber.App
	// db  *dbConn.Queries
	ctx context.Context

	PostController controller.UserController
	UserRoutes     routes.UserRouter
)

func init() {
	ini.LoadEnvVariables()
	ini.ConnecttoDB()
}

func main() {

	api := app.Group("/api")

	api.Get("/healthchecker", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"status": "success", "message": "Welcome to Golang with PostgreSQL"})
	})
}
