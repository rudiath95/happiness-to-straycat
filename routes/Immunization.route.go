package routes

import (
	controllers "happiness-to-straycat/controller"
	db "happiness-to-straycat/db/sqlc"
	"happiness-to-straycat/middleware"

	"github.com/gofiber/fiber/v2"
)

type ImmunizationRoutes struct {
	immunizationController controllers.ImmunizationController
	db                     *db.Queries
}

func NewImmunizationRoutes(immunizationController controllers.ImmunizationController, db *db.Queries) ImmunizationRoutes {
	return ImmunizationRoutes{immunizationController, db}
}

func (rc *ImmunizationRoutes) ImmunizationRoute(rg fiber.Router) {

	router := rg.Group("/immunization")
	router.Post("/", middleware.DeserializeUser(rc.db), rc.immunizationController.CreateImmunization)
	router.Get("/", middleware.DeserializeUser(rc.db), rc.immunizationController.GetAllImmunizations)
	router.Patch("/:immuneId", middleware.DeserializeUser(rc.db), rc.immunizationController.UpdateImmunization)
	router.Get("/:immuneId", middleware.DeserializeUser(rc.db), rc.immunizationController.GetImmunizationById)
	router.Delete("/:immuneId", middleware.DeserializeUser(rc.db), rc.immunizationController.DeleteImmunizationById)
}
