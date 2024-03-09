package controllers

import (
	"github.com/NelsonelTech/Nelsoneltech-Fiber/database"
	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/flash"
	"github.com/NelsonelTech/Nelsoneltech-Fiber/models"
)

func Home(c *fiber.Ctx) error{
	return c.Render("home", c, "partials/layout")
}

func About(c *fiber.Ctx) error{
	return c.Render("about", c, "partials/layout")
}

func Services(c *fiber.Ctx) error{
	successMessage := flash.Get(c)
	// Get the database connection
	db := database.DBConn 

	// Fetch all services with their associated images from the database
	var services []models.Service
	if err := db.Find(&services).Error; err != nil {
		return err
	}

	context := fiber.Map{
		"Services": services,
		"successMessage": successMessage,
	}
	return c.Render("services", context, "partials/layout")
}

func Projects(c *fiber.Ctx) error{
	return c.Render("projects", c, "partials/layout")
}

func Blogs(c *fiber.Ctx) error{
	return c.Render("blogs", c, "partials/layout")
}

func Contact(c *fiber.Ctx) error{
	return c.Render("contact", c, "partials/layout")
}