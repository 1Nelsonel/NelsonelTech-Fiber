package admincontrollers

import "github.com/gofiber/fiber/v2"

func Dashboard(c *fiber.Ctx) error {
    return c.Render("dashboard", c, "partials/adminLayout")    
}

// 