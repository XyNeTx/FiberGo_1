package main

import (
	"encoding/json"
	"fiberGo_1/internal/model"
	"fiberGo_1/internal/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		c.Accepts("json", "text")
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/Employees/:EmployeeID", func(c *fiber.Ctx) error {
		var emp model.Employee
		id := c.Params("EmployeeID")
		emp, _ = service.GetEmployeeByID(id)
		if emp == (model.Employee{}) {
			return c.Status(404).JSON(fiber.Map{
				"status": 404,
				"error":  "Employee not found",
				"test":   "test121",
			})
		}
		return c.JSON(emp)
	})

	app.Post("/Employees", func(c *fiber.Ctx) error {
		var emp model.Employee
		if err := c.BodyParser(&emp); err != nil {
			return err
		}
		bytes, _ := json.Marshal(emp)
		fmt.Println(string(bytes))
		return c.JSON(emp)
	})

	app.Listen(":3000")
}
