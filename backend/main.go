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
		user_name := ""

		row := db.QueryRow("SELECT user_name FROM users WHERE user_id = $1", 1)
		err := row.Scan(&user_name)
		if err != nil {
			return err
		}
		return c.SendString("User ID: " + user_name)
	})

	app.Listen(":3000")

}
