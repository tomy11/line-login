package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/tomy11/line-api/repository"
	"gopkg.in/mgo.v2/bson"
)

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
	return ctx.Status(http.StatusCreated).JSON(bson.M{})
}

func (c *slipUserController) GetSlip(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(bson.M{})
}

func (c *slipUserController) GetAll(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(bson.M{})
}

func (c *slipUserController) PutSlip(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(bson.M{})
}

func (c *slipUserController) DeleteSlip(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusNoContent).JSON(bson.M{})
}
