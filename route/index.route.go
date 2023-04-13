package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kauri646/go-restapi-fiber/config"
	"github.com/kauri646/go-restapi-fiber/handler"
	"github.com/kauri646/go-restapi-fiber/middleware"
)



func RouteInit(r *fiber.App){

	r.Static("/public", config.ProjectRootPath+"/public/asset")

	r.Post("/login", handler.LoginHandler)
	r.Get("/user", middleware.Auth, handler.UserHandlerGetAll)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Post("/user", handler.UserHandlerCreate)
	r.Put("/user/:id", handler.UserHandlerUpdate)
	r.Put("/user/:id/update-email", handler.UserHandlerUpdateEmail)
	r.Delete("/user/:id", handler.UserHandlerDelete)

	r.Get("/theater", handler.TheaterHandlerGetAll)
	r.Get("/theater/:kota", handler.TheaterHandlerGetById)
	r.Get("/theaterdetails/details", handler.TheaterHandlerGetDetails)
	r.Post("/theater/lists", handler.TheaterHandlerCreateList)
	r.Post("/theater", handler.TheaterHandlerCreate)
	r.Put("/theater/:kota", handler.TheaterHandlerUpdate) 
	r.Delete("/theater/:kota", handler.TheaterHandlerDelete)

	r.Get("/films", handler.FilmHandlerGetAll)
	r.Get("/film/:id", handler.FilmHandlerGetById)
	r.Get("/films/:theaterid", handler.FilmHandlerGetByTheaterId)
	r.Post("/film", handler.FilmHandlerCreate)
	r.Put("/film/:id", handler.FilmHandlerUpdate)
	r.Delete("/film/:id", handler.FilmHandlerDelete)

	
}