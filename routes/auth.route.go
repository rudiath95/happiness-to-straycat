package routes

import (
	controllers "happiness-to-straycat/controller"
	db "happiness-to-straycat/db/sqlc"
	"happiness-to-straycat/middleware"

	"github.com/gofiber/fiber/v2"
)

type AuthRoutes struct {
	authController controllers.AuthController
	db             *db.Queries
}

func NewAuthRoutes(authController controllers.AuthController, db *db.Queries) AuthRoutes {
	return AuthRoutes{authController, db}
}

func (rc *AuthRoutes) AuthRoute(rg fiber.Router) {

	router := rg.Group("/auth")
	router.Post("/register", rc.authController.SignUpUser)
	router.Post("/login", rc.authController.SignInUser)
	router.Get("/refresh", rc.authController.RefreshAccessToken)
	router.Get("/logout", middleware.DeserializeUser(rc.db), rc.authController.LogoutUser)
}
