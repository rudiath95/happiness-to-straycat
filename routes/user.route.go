package routes

import (
	controllers "happiness-to-straycat/controller"
	db "happiness-to-straycat/db/sqlc"
	"happiness-to-straycat/middleware"

	"github.com/gofiber/fiber/v2"
)

type UserRoutes struct {
	userController controllers.UserController
	db             *db.Queries
}

func NewUserRoutes(userController controllers.UserController, db *db.Queries) UserRoutes {
	return UserRoutes{userController, db}
}

func (rc *UserRoutes) UserRoute(rg fiber.Router) {

	router := rg.Group("/users")
	router.Get("/me", middleware.DeserializeUser(rc.db), rc.userController.GetMe)
}
