package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v3"
	_ "github.com/lib/pq"
)

func main() {
	// Database connection
	connStr := "host=db port=5432 user=app password=password dbname=app_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		var greeting string
		err := db.QueryRow("SELECT 'Hello, World!'").Scan(&greeting)
		if err != nil {
			return err
		}
		return c.SendString(greeting)
	})

	app.Listen(":3000")

}
