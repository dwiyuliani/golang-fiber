package migration

import (
	"fmt"
	"go-fiber/database"
	"go-fiber/model/entity"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database migrated")
}
