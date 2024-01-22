package main

import (
	"go-fiber/database"
	"go-fiber/database/migration"
	"go-fiber/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//INITIAL DATABASE
	database.DatabaseInit()

	//migrate tabel seperti laravel
	migration.RunMigration()

	app := fiber.New()

	//initial route
	route.RouteInit(app)
	app.Listen(":8080")
}
