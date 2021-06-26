package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomy11/line-api/controller"
)

type pointsRoutes struct {
	pointsController controller.PointsController
}

func NewpointsRoutes(pointsController controller.PointsController) Routes {
	return &pointsRoutes{pointsController}
}

func (r *pointsRoutes) Install(app *fiber.App) {
	app.Post("/points", AuthRequired, r.pointsController.PostPoints)
	app.Get("/points", AuthRequired, r.pointsController.GetAllPoints)
	app.Get("/points/:id", AuthRequired, r.pointsController.GetPoints)
	app.Put("/points/:id", AuthRequired, r.pointsController.PutPoints)
	app.Delete("/points/:id", AuthRequired, r.pointsController.DeletePoints)
}
