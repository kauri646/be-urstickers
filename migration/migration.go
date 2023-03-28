package migration

import (
	"fmt"
	"log"

	"github.com/kauri646/go-restapi-fiber/database"
	"github.com/kauri646/go-restapi-fiber/model/entity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}