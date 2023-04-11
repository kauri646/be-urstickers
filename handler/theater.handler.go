package handler

import (
	"log"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/kauri646/go-restapi-fiber/database"
	"github.com/kauri646/go-restapi-fiber/model/entity"
	"github.com/kauri646/go-restapi-fiber/request"
)

func TheaterHandlerGetAll(ctx *fiber.Ctx) error {

	var theater []entity.Theater

	result := database.DB.Find(&theater)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(theater)

}

func TheaterHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("kota")
	
    var theater entity.Theater
	err := database.DB.First(&theater, "kota = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "theater not found",
		})
	}
    
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": theater,
	})
}

func TheaterHandlerCreate(ctx *fiber.Ctx) error {
	theater := new(request.TheaterCreateRequest)

	if err := ctx.BodyParser(theater); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(theater)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error": errValidate.Error(),
		})
	}


	newTheater := entity.Theater{
		Kota: theater.Kota,
        Theater: theater.Theater,
        Phone: theater.Phone,
		
	}

		
	errCreateTheater := database.DB.Create(&newTheater).Error
	if errCreateTheater != nil{
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": newTheater,
	})
}

func TheaterHandlerUpdate(ctx *fiber.Ctx) error {
	
	theaterRequest := new(request.TheaterUpdateRequest)
	if err := ctx.BodyParser(theaterRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var theater entity.Theater

	userId := ctx.Params("kota")
	err := database.DB.First(&theater, "kota = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// if theaterRequest.Kota != "" {
	// theater.Kota = theaterRequest.Kota
	// } 

	theater.Kota = theaterRequest.Kota
	theater.Theater = theaterRequest.Theater
	theater.Phone = theaterRequest.Phone

	errUpdate := database.DB.Save(&theater).Error
	if errUpdate!= nil {
        return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
			
        })
    }

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": theater,
	})
}

func TheaterHandlerDelete(ctx *fiber.Ctx) error {
	

	userId := ctx.Params("kota")

    var theater entity.Theater

	err := database.DB.Debug().First(&theater, "kota = ?", userId).Error
    if err!= nil {
        return ctx.Status(404).JSON(fiber.Map{
        	"message": "theater not found",
        })
    }
	errDelete := database.DB.Debug().Delete(&theater).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
        	"message": "internal server error",
        })
	}

	return ctx.JSON(fiber.Map{
		"message": "theater was deleted",
	})
}
	