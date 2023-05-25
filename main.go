package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/AshirwadPradhan/hotelresapi/api"
	"github.com/AshirwadPradhan/hotelresapi/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DBURI = "mongodb://localhost:27017"
const DBNAME = "hotel_reservation"
const USERCOLL = "users"

var errConf = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	}}

func main() {
	fmt.Println("Hotel Reservation API")
	listenaddr := flag.String("listenAddr", ":5000", "The listen address of API server")
	flag.Parse()

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DBURI))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	userHandler := api.NewUserHandler(db.NewMongoStore(client))

	app := fiber.New(errConf)
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/users", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	apiv1.Post("/users", userHandler.HandlePostUser)

	app.Listen(*listenaddr)
}
