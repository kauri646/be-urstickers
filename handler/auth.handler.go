package handler

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/kauri646/go-restapi-fiber/database"
	"github.com/kauri646/go-restapi-fiber/internal/models/users/entity"
	"github.com/kauri646/go-restapi-fiber/request"
	"github.com/kauri646/go-restapi-fiber/utils"
)

func LoginHandler(ctx *fiber.Ctx) error {
	LoginRequest := new(request.LoginRequest)

	if err := ctx.BodyParser(LoginRequest); err != nil {
		return err
	}
	log.Println(LoginRequest)

	validate := validator.New()
	errValidate := validate.Struct(LoginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var user entity.User
	err := database.DB.First(&user, "email = ?", LoginRequest.Email).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	isValid := utils.CheckPasswordHash(LoginRequest.Password, user.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credentials",
		})
	}

	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["address"] = user.Address
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	if user.Email == "atra@gmail.com" {
		claims["role"] = "admin"
	} else {
		claims["role"] = "user"
	}

	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credentials",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}
