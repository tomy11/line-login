package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomy11/line-api/controller"
)

type pointToProductRoutes struct {
	pointToProductController controller.PointToProductController
}

func NewPointToProductRoutes(pointToProductController controller.PointToProductController) Routes {
	return &pointToProductRoutes{pointToProductController}
}

func (r *pointToProductRoutes) Install(app *fiber.App) {
	app.Post("/pointToProduct", AuthRequired, r.pointToProductController.PostPointToProduct)
	app.Get("/pointToProduct", AuthRequired, r.pointToProductController.GetAllPointToProduct)
	app.Get("/pointToProduct/:id", AuthRequired, r.pointToProductController.GetAllPointToProduct)
	app.Put("/pointToProduct/:id", AuthRequired, r.pointToProductController.PutPointToProduct)
	app.Delete("/pointToProduct/:id", AuthRequired, r.pointToProductController.DeletePointToProduct)
}
