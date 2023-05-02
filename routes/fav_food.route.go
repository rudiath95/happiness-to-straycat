package routes

import (
	controllers "happiness-to-straycat/controller"
	db "happiness-to-straycat/db/sqlc"
	"happiness-to-straycat/middleware"

	"github.com/gofiber/fiber/v2"
)

type FavFoodRoutes struct {
	foodController controllers.FavFoodController
	db             *db.Queries
}

func NewFavFoodRoutes(favFoodController controllers.FavFoodController, db *db.Queries) FavFoodRoutes {
	return FavFoodRoutes{favFoodController, db}
}

func (rc *FavFoodRoutes) FavFoodRoute(rg fiber.Router) {

	router := rg.Group("/food")
	router.Post("/", middleware.DeserializeUser(rc.db), rc.foodController.CreateFavFood)
	router.Get("/", middleware.DeserializeUser(rc.db), rc.foodController.GetAllFavFoods)
	router.Patch("/:foodId", middleware.DeserializeUser(rc.db), rc.foodController.UpdateFavFood)
	router.Get("/:foodId", middleware.DeserializeUser(rc.db), rc.foodController.GetFavFoodById)
	router.Delete("/:foodId", middleware.DeserializeUser(rc.db), rc.foodController.DeleteFavFoodById)
}
