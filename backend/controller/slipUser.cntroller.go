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

type CommentHandler struct {
	PointColl *mgo.Collection
	SlipColl  *mgo.Collation
}
type SlipUserController interface {
	PostSlip(ctx *fiber.Ctx) error
	GetSlip(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	PutSlip(ctx *fiber.Ctx) error
	DeleteSlip(ctx *fiber.Ctx) error
}

type slipUserController struct {
	slipUserRepo repository.SlipUserRepository
}

func NewSlipUserController(slipUserRepo repository.SlipUserRepository) SlipUserController {
	return &slipUserController{slipUserRepo}
}

func (c *slipUserController) PostSlip(ctx *fiber.Ctx) error {
	var inputSlip models.InputSlipUser

	err := ctx.BodyParser(&inputSlip)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}

	token := ctx.Locals("user").(*jwt.Token)
	payload, err := security.ParseToken(token.Raw)
	if err != nil {
		panic(err)
	}
	sid := bson.NewObjectId()
	slip := models.SlipUser{
		Id:        sid,
		LineId:    inputSlip.LineId,
		Images:    inputSlip.Images,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserId:    models.Auth{UserId: payload.Id},
	}

	err = c.slipUserRepo.Save(&slip)
	if err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}

	return ctx.Status(http.StatusCreated).JSON(bson.M{"Message": "save data successfuly "})
}

func NewAuth(userId string) models.Auth {
	return models.Auth{
		UserId: userId,
	}
}

func (c *slipUserController) GetSlip(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		panic("not find id")
	}
	token := ctx.Locals("user").(*jwt.Token)
	payload, err := security.ParseToken(token.Raw)
	if err != nil {
		panic(err)
	}
	terget, err := c.slipUserRepo.GetById(id)
	if err != nil {
		panic(err)
	}

	if err == mgo.ErrNotFound || terget.UserId.UserId != payload.Id {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(terget)
}

func (c *slipUserController) GetAll(ctx *fiber.Ctx) error {
	users, err := c.slipUserRepo.GetAll()
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(users)
}

func (c *slipUserController) PutSlip(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(bson.M{})
}

func (c *slipUserController) DeleteSlip(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		panic("not find id")
	}
	token := ctx.Locals("user").(*jwt.Token)
	payload, err := security.ParseToken(token.Raw)
	if err != nil {
		panic(err)
	}
	terget, err := c.slipUserRepo.GetById(id)
	if err != nil {
		panic(err)
	}

	if err == mgo.ErrNotFound || terget.UserId.UserId != payload.Id {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}
	c.slipUserRepo.Delete(id)
	return ctx.Status(http.StatusNoContent).JSON(bson.M{"message": "delete success"})

}

func (p CommentHandler) SavePoint(userId string) {
	points := models.Point{
		UserId:    userId,
		SlipId:    "",
		Point:     0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := p.PointColl.Insert(&points)
	if err != nil {
		panic(err)
	}
}
