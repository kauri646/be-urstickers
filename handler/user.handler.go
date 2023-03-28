package handler

import (
	"log"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/kauri646/go-restapi-fiber/database"
	"github.com/kauri646/go-restapi-fiber/model/entity"
	"github.com/kauri646/go-restapi-fiber/request"
	"github.com/kauri646/go-restapi-fiber/utils"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {

	var users []entity.User
	
	result := database.DB.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(users)
	
}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error": errValidate.Error(),
		})
	}


	newUser := entity.User{
		Name: user.Name,
        Email: user.Email,
        Address: user.Address,
        Phone: user.Phone,
	}

	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "iternal server error",
            })
		
	}

	newUser.Password = hashedPassword


	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil{
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": newUser,
	})
}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	
    var user entity.User
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// userResponse := response.UserResponse{
	// 	ID: user.ID,
	// 	Name: user.Name,
	// 	Email: user.Email,
    //     Address: user.Address,
    //     Phone: user.Phone,
	// 	CreatedAt: user.CreatedAt,
	// 	UpdatedAt: user.UpdatedAt,
    // }

    
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": user,
	})
}

func UserHandlerUpdate(ctx *fiber.Ctx) error {
	
	userRequest := new(request.UserUpdateRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user entity.User

	userId := ctx.Params("id")
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if userRequest.Name != "" {
	user.Name = userRequest.Name
	} 

	user.Address = userRequest.Address
	user.Phone = userRequest.Phone
	errUpdate := database.DB.Save(&user).Error
	if errUpdate!= nil {
        return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
			
        })
    }

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": user,
	})
}

func UserHandlerUpdateEmail(ctx *fiber.Ctx) error {
	
	userRequest := new(request.UserEmailRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user entity.User
	var isEmailUserExist entity.User

	userId := ctx.Params("id")
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	errCheckEmail := database.DB.First(&isEmailUserExist, "email = ?", userRequest.Email).Error
	if errCheckEmail != nil {
		return ctx.Status(402).JSON(fiber.Map{
			"message": "user already used.",
		})
	}
	
	user.Email = userRequest.Email
	
	errUpdate := database.DB.Save(&user).Error
	if errUpdate!= nil {
        return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
			
        })
    }

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": user,
	})
}

func UserHandlerDelete(ctx *fiber.Ctx) error {
	

	userId := ctx.Params("id")

    var user entity.User

	err := database.DB.Debug().First(&user, "id=?", userId).Error
    if err!= nil {
        return ctx.Status(404).JSON(fiber.Map{
        	"message": "user not found",
        })
    }
	errDelete := database.DB.Debug().Delete(&user).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
        	"message": "internal server error",
        })
	}

	return ctx.JSON(fiber.Map{
		"message": "user was deleted",
	})
}
	

