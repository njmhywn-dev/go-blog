package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/njmhywn-dev/go-blog/controller"
	"github.com/njmhywn-dev/go-blog/middleware"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/blog", controller.BlogList)
	app.Get("/blog/:id", controller.BlogDetail)
	app.Post("/blog", controller.BlogCreate)
	app.Put("/blog/:id", controller.BlogUpdate)
	app.Delete("/blog/:id", controller.BlogDelete)

	app.Post("/login", controller.Login)
	app.Post("/register", controller.Register)

	private := app.Group("/private")
	private.Use(middleware.Authenticate)

	private.Get("/refreshtoken", controller.RefreshToken)
	private.Get("/profile", controller.Profile)

}
