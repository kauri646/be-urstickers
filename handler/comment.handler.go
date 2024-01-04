package handler

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/kauri646/go-restapi-fiber/database"
	"github.com/kauri646/go-restapi-fiber/internal/models/users/entity"
	"github.com/kauri646/go-restapi-fiber/request"
)

func CommentHandlerCreate(ctx *fiber.Ctx) error {
	comment := new(request.CommentCreateRequest)

	if err := ctx.BodyParser(comment); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(comment)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newComment := entity.Comment{
		FilmId:  comment.FilmId,
		Comment: comment.Comment,
	}

	errCreateUser := database.DB.Create(&newComment).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newComment,
	})
}
