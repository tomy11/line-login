package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomy11/line-api/controller"
)

type authRoutes struct {
	authController controller.AuthController
}

func NewAuthRoutes(authController controller.AuthController) Routes {
	return &authRoutes{authController}
}

func (r *authRoutes) Install(app *fiber.App) {
	app.Post("/signup", r.authController.SignUp)
	app.Post("/signin", r.authController.SignIn)
	app.Get("/users", AuthRequired, r.authController.GetUsers)
	app.Get("/users/:id", AuthRequired, r.authController.GetUser)
	app.Put("/users/:id", AuthRequired, r.authController.PutUser)
	app.Delete("/users/:id", AuthRequired, r.authController.DeleteUser)
}
