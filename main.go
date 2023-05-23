package main

import (
	"flag"
	"fmt"

	"github.com/AshirwadPradhan/hotelresapi/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hotel Reservation API")
	listenaddr := flag.String("listenAddr", ":5000", "The listen address of API server")
	flag.Parse()

	app := fiber.New()
	app.Get("/foo", handleFoo)

	apiv1 := app.Group("/api/v1")
	apiv1.Get("/users", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*listenaddr)
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"test": "serving for foo"})
}
