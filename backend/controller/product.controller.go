package controller

import (
	"net/http"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/tomy11/line-api/models"
	"github.com/tomy11/line-api/repository"
	"github.com/tomy11/line-api/security"
	"github.com/tomy11/line-api/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductController interface {
	PostProduct(ctx *fiber.Ctx) error
	GetProduct(ctx *fiber.Ctx) error
	GetAllProduct(ctx *fiber.Ctx) error
	PutProduct(ctx *fiber.Ctx) error
	DeleteProduct(ctx *fiber.Ctx) error
}

type productController struct {
	productRepo repository.ProductRepository
}

func NewProductController(productRepo repository.ProductRepository) ProductController {
	return &productController{productRepo}
}

func (c *productController) PostProduct(ctx *fiber.Ctx) error {
	var inputPro models.InputProduct
	err := ctx.BodyParser(&inputPro)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}

	if inputPro.ProductPoint <= 0 {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(bson.M{"message": "ProductPoint number must be greater than 0 "})
	}

	token := ctx.Locals("user").(*jwt.Token)
	payload, err := security.ParseToken(token.Raw)
	if err != nil {
		panic(err)
	}

	prod := models.Product{
		ProductName:  inputPro.ProductName,
		Image:        inputPro.Image,
		ProductPoint: inputPro.ProductPoint,
		CreateBy:     payload.Id,
		UpdateBy:     "",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = c.productRepo.Save(&prod)
	if err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}

	return ctx.Status(http.StatusCreated).JSON(bson.M{"Message": "save data successfuly"})
}

func (c *productController) GetProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		panic("not find id")
	}
	token := ctx.Locals("user").(*jwt.Token)
	payload, err := security.ParseToken(token.Raw)
	if err != nil {
		panic(err)
	}
	terget, err := c.productRepo.GetById(id)
	if err != nil {
		panic(err)
	}

	if err == mgo.ErrNotFound || terget.CreateBy != payload.Id {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(terget)
}

func (c *productController) GetAllProduct(ctx *fiber.Ctx) error {
	prod, err := c.productRepo.GetAll()
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(prod)
}

func (c *productController) PutProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		panic("not find id")
	}
	var Pro models.Product
	err := ctx.BodyParser(&Pro)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}
	if Pro.ProductPoint <= 0 {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(bson.M{"message": "ProductPoint number must be greater than 0 "})
	}
	token := ctx.Locals("user").(*jwt.Token)
	payload, err := security.ParseToken(token.Raw)
	if err != nil {
		panic(err)
	}

	prod := models.Product{
		Id:           Pro.Id,
		ProductName:  Pro.ProductName,
		Image:        Pro.Image,
		ProductPoint: Pro.ProductPoint,
		CreateBy:     Pro.CreateBy,
		UpdateBy:     payload.Id,
		CreatedAt:    Pro.CreatedAt,
		UpdatedAt:    time.Now(),
	}

	err = c.productRepo.Update(&prod)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(bson.M{"Message": "Update data successfuly"})
}

func (c *productController) DeleteProduct(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		panic("not find id")
	}
	token := ctx.Locals("user").(*jwt.Token)
	payload, err := security.ParseToken(token.Raw)
	if err != nil {
		panic(err)
	}
	terget, err := c.productRepo.GetById(id)
	if err != nil {
		panic(err)
	}

	if err == mgo.ErrNotFound || terget.CreateBy != payload.Id {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}
	c.productRepo.Delete(id)
	return ctx.Status(http.StatusNoContent).JSON(bson.M{"message": "delete success"})

}
