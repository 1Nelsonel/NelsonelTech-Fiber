package main

import (
	adminroutes "github.com/NelsonelTech/Nelsoneltech-Fiber/admin/adminRoutes"
	"github.com/NelsonelTech/Nelsoneltech-Fiber/database"
	"github.com/NelsonelTech/Nelsoneltech-Fiber/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

// Middleware to initialize db connections
func init() {
	database.ConnectDB()
}


func main() {

	// Create a new engine
	engine := html.New("./views", ".html")
	
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "partials/layouts",
	})

	app.Static("static", "./static")
	app.Static("staticadmin", "./staticadmin")

	// Initialize default config
	app.Use(logger.New())

	app.Use(cors.New())

	// Routes
	routes.SetupRoutes(app)
	adminroutes.AdminSetupRoutes(app)

	// app.Listen(":8080")
	app.Listen(":8080")
}