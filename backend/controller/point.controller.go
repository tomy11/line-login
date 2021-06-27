package controller

import (
	"fmt"
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

type PointsController interface {
	PostPoints(ctx *fiber.Ctx) error
	GetPoints(ctx *fiber.Ctx) error
	GetAllPoints(ctx *fiber.Ctx) error
	PutPoints(ctx *fiber.Ctx) error
	DeletePoints(ctx *fiber.Ctx) error
	GetPointByUserId(ctx *fiber.Ctx) error
}

type pointsController struct {
	pointsRepo   repository.PointRepository
	slipUserRepo repository.SlipUserRepository
}

func NewPointsController(pointsRepo repository.PointRepository, slipUserRepo repository.SlipUserRepository) PointsController {
	return &pointsController{pointsRepo, slipUserRepo}
}

func (c *pointsController) PostPoints(ctx *fiber.Ctx) error {
	var inputPoints models.InputPoints
	err := ctx.BodyParser(&inputPoints)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}

	if inputPoints.Point <= 0 {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(bson.M{"message": "Point number must be greater than 0 "})
	}

	points := models.Point{
		UserId:    inputPoints.UserId,
		SlipId:    inputPoints.SlipId,
		Point:     inputPoints.Point,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = c.pointsRepo.Save(&points)
	if err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}

	return ctx.Status(http.StatusCreated).JSON(bson.M{"Message": "save data successfuly"})
}

func (c *pointsController) GetPoints(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		panic("not find id")
	}
	token := ctx.Locals("user").(*jwt.Token)
	payload, err := security.ParseToken(token.Raw)
	if err != nil {
		panic(err)
	}
	terget, err := c.pointsRepo.GetById(id)
	if err != nil {
		panic(err)
	}

	results, err := c.slipUserRepo.GetById("60d75738c70ad0408f9b1b42")
	if err != nil {
		panic(err)
	}

	fmt.Println("res :", results)

	if err == mgo.ErrNotFound || terget.UserId != payload.Id {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(terget)
}

func (c *pointsController) GetAllPoints(ctx *fiber.Ctx) error {
	pot, err := c.pointsRepo.GetAll()
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(pot)
}

func (c *pointsController) PutPoints(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		panic("not find id")
	}
	var pot models.Point
	err := ctx.BodyParser(&pot)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}
	if pot.Point <= 0 {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(bson.M{"message": "Point number must be greater than 0 "})
	}
	token := ctx.Locals("user").(*jwt.Token)
	payload, err := security.ParseToken(token.Raw)
	if err != nil {
		panic(err)
	}

	if err == mgo.ErrNotFound || pot.UserId != payload.Id {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}

	points := models.Point{
		Id:        bson.ObjectId(pot.Id.Hex()),
		UserId:    pot.UserId,
		SlipId:    pot.SlipId,
		CreatedAt: pot.CreatedAt,
		UpdatedAt: time.Now(),
	}

	err = c.pointsRepo.Update(&points)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(bson.M{"Message": "Update data successfuly"})
}

func (c *pointsController) DeletePoints(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		panic("not find id")
	}
	token := ctx.Locals("user").(*jwt.Token)
	payload, err := security.ParseToken(token.Raw)
	if err != nil {
		panic(err)
	}
	terget, err := c.pointsRepo.GetById(id)
	if err != nil {
		panic(err)
	}

	if err == mgo.ErrNotFound || terget.UserId != payload.Id {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}
	c.pointsRepo.Delete(id)
	return ctx.Status(http.StatusNoContent).JSON(bson.M{"message": "delete success"})

}

func (c *pointsController) GetPointByUserId(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		panic("not find id")
	}
	terget, err := c.pointsRepo.GetPointByUserId(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(terget)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(terget)
}
