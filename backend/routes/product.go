package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomy11/line-api/controller"
)

type productRoutes struct {
	productController controller.ProductController
}

func NewProductRoutes(productController controller.ProductController) Routes {
	return &productRoutes{productController}
}

func (r *productRoutes) Install(app *fiber.App) {
	app.Post("/product", AuthRequired, r.productController.PostProduct)
	app.Get("/product", AuthRequired, r.productController.GetAllProduct)
	app.Get("/product/:id", AuthRequired, r.productController.GetProduct)
	app.Put("/product/:id", AuthRequired, r.productController.PutProduct)
	app.Delete("/product/:id", AuthRequired, r.productController.DeleteProduct)
}
