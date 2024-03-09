package adminroutes

import (
	admincontrollers "github.com/NelsonelTech/Nelsoneltech-Fiber/admin/adminControllers"
	"github.com/gofiber/fiber/v2"
)

func AdminSetupRoutes(app *fiber.App) {
    app.Get("/dashboard/", admincontrollers.Dashboard)

	// services
	app.Get("/service/list/", admincontrollers.ListServices)  
	app.Get("service/add/", admincontrollers.ServiceAdd)
	app.Post("/service-add/", admincontrollers.AddService)
	app.Get("/service/edit/:slug/", admincontrollers.EditService)
	app.Post("/service/update/:slug/", admincontrollers.UpdateService)
	app.Post("/service/delete/:slug/", admincontrollers.DeleteService)
}