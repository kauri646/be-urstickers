package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kauri646/go-restapi-fiber/database"
	"github.com/kauri646/go-restapi-fiber/migration"
	"github.com/kauri646/go-restapi-fiber/route"
)

func main(){

	database.DatabaseInit()
	migration.RunMigration()
	
	app := fiber.New()

	route.RouteInit(app)

	app.Listen(":8080")
}