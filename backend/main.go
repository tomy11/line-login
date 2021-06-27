package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/tomy11/line-api/controller"
	"github.com/tomy11/line-api/db"
	"github.com/tomy11/line-api/repository"
	"github.com/tomy11/line-api/routes"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}
}

func main() {
	conn := db.NewConnection()
	defer conn.Close()

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Hello Word"})
	})

	usersRepo := repository.NewUsersRepository(conn)
	slipUserRepo := repository.NewSlipUserRepository(conn)
	productRepo := repository.NewProductRepository(conn)
	pointsPepo := repository.NewPointRepository(conn)
	pointToProductRepo := repository.NewPointToProductRepository(conn)

	authController := controller.NewAuthController(usersRepo)
	pointController := controller.NewPointsController(pointsPepo, slipUserRepo)
	slipUserController := controller.NewSlipUserController(slipUserRepo)
	productController := controller.NewProductController(productRepo)
	pointToProductController := controller.NewPointToProductController(pointToProductRepo)

	authRoutes := routes.NewAuthRoutes(authController)
	pointRoutes := routes.NewpointsRoutes(pointController)
	slipUserRoutes := routes.NewSlipUserRoutes(slipUserController)
	productRoutes := routes.NewProductRoutes(productController)
	pointToProductRoutes := routes.NewPointToProductRoutes(pointToProductController)

	authRoutes.Install(app)
	slipUserRoutes.Install(app)
	productRoutes.Install(app)
	pointRoutes.Install(app)
	pointToProductRoutes.Install(app)

	app.Listen(":3000")
}
