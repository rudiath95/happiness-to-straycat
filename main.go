package main

import (
	"context"
	"database/sql"
	"fmt"
	"happiness-to-straycat/config"
	controllers "happiness-to-straycat/controller"
	dbConn "happiness-to-straycat/db/sqlc"
	"happiness-to-straycat/routes"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
)

var (
	app *fiber.App
	db  *dbConn.Queries
	ctx context.Context

	AuthController controllers.AuthController
	UserController controllers.UserController
	AuthRoutes     routes.AuthRoutes
	UserRoutes     routes.UserRoutes
)

func init() {
	ctx = context.TODO()
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	conn, err := sql.Open(config.PostgreDriver, config.PostgresSource)
	if err != nil {
		log.Fatalf("could not connect to postgres database: %v", err)
	}

	db = dbConn.New(conn)

	fmt.Println("PostgreSQL connected successfully...")

	AuthController = *controllers.NewAuthController(db, ctx)
	UserController = controllers.NewUserController(db, ctx)
	AuthRoutes = routes.NewAuthRoutes(AuthController, db)
	UserRoutes = routes.NewUserRoutes(UserController, db)

	app = fiber.New()
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	corsConfig := cors.Config{
		AllowOrigins:     config.Origin,
		AllowCredentials: true,
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
	}

	app.Use(cors.New(corsConfig))

	// MANUAL CORS CONFIGURATION
	// app.Use(func(c *fiber.Ctx) error {
	// 	c.Set("Access-Control-Allow-Origin", "*") //change * to custom domain later
	// 	c.Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
	// 	c.Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

	// 	if c.Method() == "OPTIONS" {
	// 		return c.SendString("allowed")
	// 	}

	// 	return c.SendString("hello")
	// })

	router := app.Group("/api")

	router.Get("/healthchecker", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"status": "success", "message": "Welcome to Golang with PostgreSQL"})
	})

	AuthRoutes.AuthRoute(router)
	UserRoutes.UserRoute(router)

	app.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("Route %s not found", ctx.Path())})
	})

	log.Fatal(app.Listen(":" + config.Port))
}
