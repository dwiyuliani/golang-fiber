package route

import (
	"go-fiber/config"
	"go-fiber/handler"
	"go-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"/public/asset")

	r.Post("/login", handler.LoginHandler)

	r.Get("/user", middleware.Auth, handler.UserHandlerGetAll)
	r.Post("/create", handler.UserHandlerCreate)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Put("/user/:id", handler.UserHandlerUpdate)
	r.Put("/user/:id/update-email", handler.UserHandlerUpdateEmail)
	r.Delete("/user/:id", handler.UserHandlerDelete)

}
