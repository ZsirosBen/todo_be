package main

import (
	"fmt"

	"github.com/ZsirosB/todo-go-be/database"
	"github.com/ZsirosB/todo-go-be/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello world")
}

func initDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=Morta0l1! dbname=postgres port=5432"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connected")
	database.DBConn.AutoMigrate(&models.Todo{})
	fmt.Println("Migrated DB")
}

func setupRoutes(app *fiber.App) {
	app.Get("/todos", models.GetTodos)
	app.Post("/todos", models.CreateTodo)
	app.Get("/todos/:id", models.GetTodoById)
	app.Put("/todos/:id", models.UpdateTodo)
	app.Delete("/todos/:id", models.DeleteTodo)
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	initDatabase()
	setupRoutes(app)

	app.Get("/", helloWorld)
	app.Listen(":8080")
}
