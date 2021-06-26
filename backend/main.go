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
	authController := controller.NewAuthController(usersRepo)
	authRoutes := routes.NewAuthRoutes(authController)

	pointsPepo := repository.NewPointRepository(conn)
	pointController := controller.NewPointsController(pointsPepo)
	pointRoutes := routes.NewpointsRoutes(pointController)

	slipUserRepo := repository.NewSlipUserRepository(conn)
	slipUserController := controller.NewSlipUserController(slipUserRepo)
	slipUserRoutes := routes.NewSlipUserRoutes(slipUserController)

	productRepo := repository.NewProductRepository(conn)
	productController := controller.NewProductController(productRepo)
	productRoutes := routes.NewProductRoutes(productController)

	pointToProductRepo := repository.NewPointToProductRepository(conn)
	pointToProductController := controller.NewPointToProductController(pointToProductRepo)
	pointToProductRoutes := routes.NewPointToProductRoutes(pointToProductController)

	authRoutes.Install(app)
	slipUserRoutes.Install(app)
	productRoutes.Install(app)
	pointRoutes.Install(app)
	pointToProductRoutes.Install(app)

	app.Listen(":3000")
}
