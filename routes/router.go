package routes

import (
	"github.com/NelsonelTech/Nelsoneltech-Fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controllers.Home)
	app.Get("/about/", controllers.About)
	app.Get("/services/", controllers.Services)
	app.Get("/projects/", controllers.Projects)
	app.Get("/blogs/", controllers.Blogs)
	app.Get("/contact/", controllers.Contact)
}