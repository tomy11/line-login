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

type PointToProductController interface {
	PostPointToProduct(ctx *fiber.Ctx) error
	GetPointToProduct(ctx *fiber.Ctx) error
	GetAllPointToProduct(ctx *fiber.Ctx) error
	PutPointToProduct(ctx *fiber.Ctx) error
	DeletePointToProduct(ctx *fiber.Ctx) error
}

type pointToProductController struct {
	pointToProductRepo repository.PointToProductRepository
}

func NewPointToProductController(pointToProductRepo repository.PointToProductRepository) PointToProductController {
	return &pointToProductController{pointToProductRepo}
}

func (c *pointToProductController) PostPointToProduct(ctx *fiber.Ctx) error {
	var input models.InputPointToProduct
	err := ctx.BodyParser(&input)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}

	if input.Point <= 0 {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(bson.M{"message": "ProductPoint number must be greater than 0 "})
	}

	token := ctx.Locals("user").(*jwt.Token)
	payload, err := security.ParseToken(token.Raw)
	if err != nil {
		panic(err)
	}

	ptoProd := models.PointToProduct{
		ProductId: input.ProductId,
		UserId:    input.UserId,
		Point:     input.Point,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Product: models.ProductWich{
			Id:           bson.NewObjectId(),
			ProductName:  input.Product.ProductName,
			Image:        input.Product.Image,
			ProductPoint: input.Product.ProductPoint,
			CreateBy:     payload.Id,
		},
	}

	err = c.pointToProductRepo.Save(&ptoProd)
	if err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}

	return ctx.Status(http.StatusCreated).JSON(bson.M{"Message": "save data successfuly"})
}

func (c *pointToProductController) GetPointToProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		panic("not find id")
	}
	token := ctx.Locals("user").(*jwt.Token)
	payload, err := security.ParseToken(token.Raw)
	if err != nil {
		panic(err)
	}
	terget, err := c.pointToProductRepo.GetById(id)
	if err != nil {
		panic(err)
	}

	if err == mgo.ErrNotFound || terget.UserId != payload.Id {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(terget)
}

func (c *pointToProductController) GetAllPointToProduct(ctx *fiber.Ctx) error {
	prod, err := c.pointToProductRepo.GetAll()
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(prod)
}

func (c *pointToProductController) PutPointToProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		panic("not find id")
	}
	var ptp models.PointToProduct
	err := ctx.BodyParser(&ptp)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}
	if ptp.Point <= 0 {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(bson.M{"message": "ProductPoint number must be greater than 0 "})
	}
	token := ctx.Locals("user").(*jwt.Token)
	payload, err := security.ParseToken(token.Raw)
	if err != nil {
		panic(err)
	}

	ptps := models.PointToProduct{
		Id:        ptp.Id,
		ProductId: ptp.ProductId,
		UserId:    ptp.UserId,
		Point:     ptp.Point,
		CreatedAt: ptp.CreatedAt,
		UpdatedAt: time.Now(),
		Product: models.ProductWich{
			Id:           bson.NewObjectId(),
			ProductName:  ptp.Product.ProductName,
			Image:        ptp.Product.Image,
			ProductPoint: ptp.Product.ProductPoint,
			CreateBy:     payload.Id,
		},
	}

	err = c.pointToProductRepo.Update(&ptps)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(bson.M{"Message": "Update data successfuly"})
}

func (c *pointToProductController) DeletePointToProduct(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		panic("not find id")
	}
	token := ctx.Locals("user").(*jwt.Token)
	payload, err := security.ParseToken(token.Raw)
	if err != nil {
		panic(err)
	}
	terget, err := c.pointToProductRepo.GetById(id)
	if err != nil {
		panic(err)
	}

	if err == mgo.ErrNotFound || terget.UserId != payload.Id {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}
	c.pointToProductRepo.Delete(id)
	return ctx.Status(http.StatusNoContent).JSON(bson.M{"message": "delete success"})

}
