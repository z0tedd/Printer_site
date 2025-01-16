package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

var db *gorm.DB

type Order struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	FilePath string `json:"file_path"`
	Total    int    `json:"total"`
}

func main() {
	// Создание сервера Fiber
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("index.html")
	})
	app.Get("/faq.html", func(c *fiber.Ctx) error {
		return c.SendFile("faq.html")
	})
	// app.Get("/images/frame.png", func(c *fiber.Ctx) error {
	// 	return c.SendFile("images/frame.png")
	// })
	app.Get("/images/qr1.png", func(c *fiber.Ctx) error {
		return c.SendFile("images/qr1.png")
	})
	app.Get("/images/qr2.png", func(c *fiber.Ctx) error {
		return c.SendFile("images/qr2.png")
	})

	// Запуск сервера
	app.Listen(":8080")
}
