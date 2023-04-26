package main

import (
	"database/sql"
	"fmt"
	"happiness-to-straycat/config"
	controllers "happiness-to-straycat/controller"
	dbConn "happiness-to-straycat/db/sqlc"
	"happiness-to-straycat/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

var (
	server *fiber.App
	db     *dbConn.Queries

	AuthController controllers.AuthController
	AuthRoutes     routes.AuthRoutes
)

func init() {
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

	AuthController = *controllers.NewAuthController(db)
	AuthRoutes = routes.NewAuthRoutes(AuthController)

	server = fiber.New()
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	router := server.Group("/api")

	router.Get("/healthchecker", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"status": "success", "message": "Welcome to Golang with PostgreSQL"})
	})

	AuthRoutes.AuthRoute(router)
	log.Fatal(server.Listen(":" + config.Port))
}
