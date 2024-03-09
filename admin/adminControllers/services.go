package admincontrollers

import (
	"fmt"

	"github.com/NelsonelTech/Nelsoneltech-Fiber/database"
	"github.com/NelsonelTech/Nelsoneltech-Fiber/models"
	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/flash"
	"strconv"
)

func ListServices(c *fiber.Ctx) error {
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

	return c.Render("list_services", context, "partials/adminLayout")
}

func ServiceAdd(c *fiber.Ctx) error {
	successMessage := flash.Get(c)

	context := fiber.Map{
		"successMessage": successMessage,
	}
    return c.Render("add_service", context, "partials/adminLayout")    
}

// Add Service
func AddService(c *fiber.Ctx) error {
	db := database.DBConn
	// Get Service name from form
	serviceName := c.FormValue("name")
	servicePrice := c.FormValue("price")
	serviceDescription := c.FormValue("description")

	// Get Service image from form
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	// Save Service image
	filePath := fmt.Sprintf("/media/uploads/service%s", file.Filename)

	// Save Service image
	if err := c.SaveFile(file, "."+filePath); err != nil {
		return err
	}

	price, err := strconv.ParseFloat(servicePrice, 64)
	if err != nil {
		// Handle the error (e.g., log it)
		return err
	}

	// Create a new Service instance
	newService := models.Service{
		Name:        serviceName,
		Price:       price,
		Description: serviceDescription,
		ServiceImage:filePath,
	}
	

	// Save the Service to the database
	if err := db.Create(&newService).Error; err != nil {
		return err
	}

	mp := fiber.Map{
		"success": true,
		"tag":"alert-sucess",
		"message": "Service added succesfully...",
	}

	// Redirect or respond as needed
	return flash.WithSuccess(c, mp).Redirect("/service/add/")
}

// EditService renders the form with existing service details
func EditService(c *fiber.Ctx) error {
    db := database.DBConn

    // Get the service slug from the request parameters
    serviceSlug := c.Params("Slug")

    // Fetch the existing service from the database using the slug
    var existingService models.Service
    if err := db.Where("Slug = ?", serviceSlug).First(&existingService).Error; err != nil {
        return err
    }

    // Render the form with the existing service data
    context := fiber.Map{
        "Service": existingService,
    }

    return c.Render("edit_service", context, "partials/adminLayout")
}


// Update Service
func UpdateService(c *fiber.Ctx) error {
	db := database.DBConn

	// Get the service ID from the request parameters
	serviceSlug := c.Params("Slug")
	
	// Fetch the existing service from the database
	var existingService models.Service
	if err := db.Where("Slug = ?", serviceSlug).First(&existingService).Error; err != nil {
        return err
    }

	// Get updated service details from the form
	serviceName := c.FormValue("name")
	servicePrice := c.FormValue("price")
	serviceDescription := c.FormValue("description")
	serviceStatus := c.FormValue("status")

	// Update service details
	existingService.Name = serviceName
	price_new, err := strconv.ParseFloat(servicePrice, 64)
	if err != nil {
		return err
	}

	existingService.Price = price_new
	existingService.Description = serviceDescription
	
	// Update the Status field
	if serviceStatus == "on" {
		existingService.Status = true
	} else if serviceStatus == "" {
		existingService.Status = false
	} else if status, err := strconv.ParseBool(serviceStatus); err == nil {
		existingService.Status = status
	} else {
		return err
	}

	// Handle image upload if a new image is provided
	file, err := c.FormFile("image")
	if err == nil {
		// Save Service image
		filePath := fmt.Sprintf("/media/uploads/service%s", file.Filename)
		if err := c.SaveFile(file, "."+filePath); err != nil {
			return err
		}
		existingService.ServiceImage = filePath
	}

	// Save the updated service to the database
	if err := db.Save(&existingService).Error; err != nil {
		return err
	}

	// If you want to update the image as well, implement the logic here

	mp := fiber.Map{
		"success": true,
		"tag":     "alert-success",
		"message": "Service updated successfully...",
	}

	// Redirect or respond as needed
	return flash.WithSuccess(c, mp).Redirect("/service/list/")
}


// DeleteService deletes a service based on its slug
func DeleteService(c *fiber.Ctx) error {
    db := database.DBConn

    // Get the service slug from the request parameters
    serviceSlug := c.Params("slug")

    // Fetch the existing service from the database
    var existingService models.Service
    if err := db.Where("slug = ?", serviceSlug).First(&existingService).Error; err != nil {
        return err
    }

    // Delete the service
    if err := db.Delete(&existingService).Error; err != nil {
        return err
    }

    // Redirect or respond as needed
    mp := fiber.Map{
        "success": true,
        "tag":     "alert-info",
        "message": "Service deleted successfully...",
    }

    return flash.WithSuccess(c, mp).Redirect("/service/list/")
}