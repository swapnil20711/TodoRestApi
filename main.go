package main

import (
	"todo-list-api/database"
	"todo-list-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	list := app.Group("/api")

	dbErr := database.InitDB()

	if dbErr != nil {
		panic(dbErr)
	}

	list.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Todo-List-API Tutorial :)")
	})

	list.Get("/tasks", routes.GetAllTasks)
	list.Get("/task/:id", routes.GetTask)
	list.Post("/add_task", routes.AddTask)
	list.Delete("/delete_task/:id", routes.DeleteTask)
	list.Patch("/update_task/:id", routes.UpdateTask)

	app.Listen(":8080")
}
