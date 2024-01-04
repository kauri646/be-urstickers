package handler

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kauri646/go-restapi-fiber/database"
	"github.com/kauri646/go-restapi-fiber/internal/models/users/entity"
	"github.com/kauri646/go-restapi-fiber/request"
	"github.com/kauri646/go-restapi-fiber/utils"
)

func FilmHandlerGetAll(ctx *fiber.Ctx) error {
	fmt.Println("ASDFGH")
	var film []entity.Film
	sql := "SELECT * FROM films"
	paging := utils.PaginationResponse{}
	// if sort := ctx.Query("sort"); sort != "" {
	// 	sql = fmt.Sprintf("%s ORDER BY judul %s", sql, sort)
	// }

	// page,_ := strconv.Atoi(ctx.Query("page", "1"))
	// 	perPage := 9
	// 	var total int64

	// 	database.DB.Raw(sql).Count(&total)

	// 	sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, perPage, (page - 1) * perPage)
	database.DB.Raw(sql).Scan(&film)

	return ctx.JSON(fiber.Map{
		"film":   film,
		"paging": paging,
	})

}

func FilmHandlerGetByTheaterId(ctx *fiber.Ctx) error {
	theaterid := ctx.QueryInt("theaterid")
	fmt.Println(theaterid)
	var film []entity.TheaterId
	err := database.DB.Raw(`
		SELECT f.id, f.judul, l.theater_id AS theater_id, f.jenis_film, f. produser, f.sutradara, f.penulis, f.produksi, f.casts, f.sinopsis, f.like
		FROM films f
		INNER JOIN lists l ON l.film_id = f.id
		WHERE l.theater_id = ?
	`, theaterid).Scan(&film)

	if err.Error != nil {
		log.Println(err.Error)
	}
	return ctx.JSON(film)
}

func FilmHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var film entity.Film
	err := database.DB.First(&film, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "film not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    film,
	})
}

func FilmHandlerCreate(ctx *fiber.Ctx) error {
	film := new(request.FilmCreateRequest)

	if err := ctx.BodyParser(film); err != nil {
		return err
	}

	file, errFile := ctx.FormFile("thumbnail")
	if errFile != nil {
		log.Println("Error File = ", errFile)
	}

	filename := file.Filename
	errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/asset/%s", filename))
	if errSaveFile != nil {
		log.Println("Fail to store file into public/thumbnails directory.")
	}

	newFilm := entity.Film{
		Judul:     film.Judul,
		Thumbnail: filename,
		JenisFilm: film.JenisFilm,
		Produser:  film.Produser,
		Sutradara: film.Sutradara,
		Penulis:   film.Penulis,
		Produksi:  film.Produksi,
		Casts:     film.Casts,
		Sinopsis:  film.Sinopsis,
		Like:      film.Like,
	}

	errCreateFilm := database.DB.Create(&newFilm).Error
	if errCreateFilm != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newFilm,
	})
}

func FilmHandlerUpdate(ctx *fiber.Ctx) error {

	filmRequest := new(request.FilmUpdateRequest)
	if err := ctx.BodyParser(filmRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var film entity.Film

	userId := ctx.Params("id")
	err := database.DB.First(&film, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "film not found",
		})
	}

	// if filmRequest.Kota != "" {
	// film.Kota = filmRequest.Kota
	// }

	/*
		Judul     string `json:"judul" validate:"required"`
		JenisFilm string `json:"jenisfilm" validate:"required"`
		Produser  string `json:"produser" validate:"required"`
		Sutradara string `json:"sutradara" validate:"required"`
		Penulis   string `json:"penulis" validate:"required"`
		Produksi  string `json:"produksi" validate:"required"`
		Casts     string `json:"casts" validate:"required"`
		Sinopsis  string `json:"sinopsis" validate:"required"`
		Like      uint   `json:"like" validate:"required"`
		Comment   string `json:"comment" validate:"required"`
	*/
	film.Judul = filmRequest.Judul
	film.JenisFilm = filmRequest.JenisFilm
	film.Produser = filmRequest.Produser
	film.Sutradara = filmRequest.Sutradara
	film.Penulis = filmRequest.Penulis
	film.Produksi = filmRequest.Produksi
	film.Casts = filmRequest.Casts
	film.Sinopsis = filmRequest.Sinopsis
	film.Like = filmRequest.Like

	errUpdate := database.DB.Save(&film).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    film,
	})
}

func FilmHandlerDelete(ctx *fiber.Ctx) error {

	userId := ctx.Params("id")

	var film entity.Film

	err := database.DB.Debug().First(&film, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "film not found",
		})
	}
	errDelete := database.DB.Debug().Delete(&film).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "film was deleted",
	})
}
