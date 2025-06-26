package main

import (
	"fiberGo_1/internal/model"
	"fiberGo_1/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/Employees/:EmployeeID", func(c *fiber.Ctx) error {
		var emp model.Employee
		id := c.Params("EmployeeID")
		emp, _ = service.GetEmployeeByID(id)
		if emp == (model.Employee{}) {
			return c.Status(404).JSON(fiber.Map{"error": "Employee not found"})
		}
		return c.JSON(emp)
	})

	app.Listen(":3000")
}
