package routes

import (
	"happiness-to-straycat/controller"

	"github.com/gofiber/fiber/v2"
)

type UserRouter struct {
	UserController *controller.UserController
}

func NewUserRouter(userController *controller.UserController) *UserRouter {
	return &UserRouter{userController}
}

func (ur *UserRouter) UserRoutes(rg *fiber.Group) {
	router := rg.Group("users")
	router.Post("/", ur.UserController.CreateUser)
}
