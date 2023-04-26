package routes

import (
	controllers "happiness-to-straycat/controller"

	"github.com/gofiber/fiber/v2"
)

type AuthRoutes struct {
	authController controllers.AuthController
}

func NewAuthRoutes(authController controllers.AuthController) AuthRoutes {
	return AuthRoutes{authController}
}

func (rc *AuthRoutes) AuthRoute(rg fiber.Router) {

	router := rg.Group("/auth")
	router.Post("/register", rc.authController.SignUpUser)
	router.Post("/login", rc.authController.SignInUser)
	// router.Get("/refresh", rc.authController.RefreshAccessToken)
	// router.Get("/logout", middleware.DeserializeUser(rc.db), rc.authController.LogoutUser)
}
