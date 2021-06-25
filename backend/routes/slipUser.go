package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomy11/line-api/controller"
)

type slipUserRoutes struct {
	slipUserController controller.SlipUserController
}

func NewSlipUserRoutes(slipUserController controller.SlipUserController) Routes {
	return &slipUserRoutes{slipUserController}
}

func (r *slipUserRoutes) Install(app *fiber.App) {
	app.Post("/slip", AuthRequired, r.slipUserController.PostSlip)
	app.Get("/slip", AuthRequired, r.slipUserController.GetAll)
	app.Get("/slip/:id", AuthRequired, r.slipUserController.GetSlip)
	app.Put("/slip/:id", AuthRequired, r.slipUserController.PutSlip)
	app.Delete("/slip/:id", AuthRequired, r.slipUserController.DeleteSlip)
}
