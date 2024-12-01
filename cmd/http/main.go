package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/pangolin-do-golang/tech-challenge-customer-api/docs"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/adapters/db/repositories"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/adapters/rest/server"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/core/customer"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// @title Tech Challenge Customer Food API
// @version 0.1.0
// @description Fast Food API for FIAP Tech course

// @host localhost:8080
// @BasePath /
func main() {
	_ = godotenv.Load()

	db, err := initDb()
	if err != nil {
		panic(err)
	}

	customerRepository := repositories.NewMongoCustomerRepository(db.Collection("customer"))
	customerService := customer.NewService(customerRepository)

	restServer := server.NewRestServer(&server.RestServerOptions{
		CustomerService: customerService,
	})

	restServer.Serve()
}

func initDb() (*mongo.Database, error) {
	dns := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	clientOpts := options.Client().ApplyURI(dns)
	client, err := mongo.Connect(clientOpts)

	if err != nil {
		log.Fatalln(err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatalln(err)
	}

	return client.Database(os.Getenv("DB_NAME")), nil
}
