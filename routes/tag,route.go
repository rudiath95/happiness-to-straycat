package routes

import (
	controllers "happiness-to-straycat/controller"
	db "happiness-to-straycat/db/sqlc"
	"happiness-to-straycat/middleware"

	"github.com/gofiber/fiber/v2"
)

type TagRoutes struct {
	tagController controllers.TagController
	db            *db.Queries
}

func NewTagRoutes(tagController controllers.TagController, db *db.Queries) TagRoutes {
	return TagRoutes{tagController, db}
}

func (rc *TagRoutes) TagRoute(rg fiber.Router) {

	router := rg.Group("/tag")
	router.Post("/", middleware.DeserializeUser(rc.db), rc.tagController.CreateTag)
	router.Get("/", middleware.DeserializeUser(rc.db), rc.tagController.GetAllTags)
	router.Patch("/:tagId", middleware.DeserializeUser(rc.db), rc.tagController.UpdateTag)
	router.Get("/:tagId", middleware.DeserializeUser(rc.db), rc.tagController.GetTagById)
	router.Delete("/:tagId", middleware.DeserializeUser(rc.db), rc.tagController.DeleteTagById)
}
