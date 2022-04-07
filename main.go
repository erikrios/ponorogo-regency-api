package main

import (
	"fmt"
	"log"
	"os"

	"github.com/erikrios/ponorogo-regency-api/config"
	"github.com/erikrios/ponorogo-regency-api/controller"
	"github.com/erikrios/ponorogo-regency-api/middleware"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %s\n", err.Error())
	}

	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Printf("Successfully connected to database with instance address: %p", db)
	}

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	homeController := controller.NewHomeController()

	e := echo.New()

	if os.Getenv("ENV") == "production" {
		middleware.BodyLimit(e)
		middleware.Gzip(e)
		middleware.RateLimiter(e)
		middleware.Recover(e)
		middleware.Secure(e)
	} else {
		middleware.Logger(e)
	}

	homeController.Route(e)

	e.Logger.Fatal(e.Start(port))
}
