package main

import (
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	app := fiber.New()
	// Add Json data
	books = append(books, Book{ID: 1, Title: "The history of me", Author: "Thananon"})
	books = append(books, Book{ID: 2, Title: "Helloworld", Author: "Pie"})

	// Method Get
	app.Get("/books", getbooks)
	app.Get("/books/:id", getbook)
	app.Post("/books", createbook)
	app.Put("/books/:id", updatebook)
	app.Delete("/books/:id", deleteBook)
	// Run at port :3000
	app.Listen(":3000")
}
