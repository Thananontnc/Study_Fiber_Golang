package main

import (
	"strconv"

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
	// Run at port :3000
	app.Listen(":3000")
}

func getbooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getbook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for _, book := range books {
		if book.ID == bookId {
			return c.JSON(book)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func createbook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	books = append(books, *book)
	return c.JSON(book)
}

func updatebook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	bookUpdate := new(Book)
	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for i, book := range books {
		if book.ID == bookId {
			books[i].Title = bookUpdate.Title
			books[i].Author = bookUpdate.Author
			return c.JSON(book)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}
