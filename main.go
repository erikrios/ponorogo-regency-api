package main

import (
	"fmt"
	"log"
	"os"

	"github.com/erikrios/ponorogo-regency-api/config"
	"github.com/erikrios/ponorogo-regency-api/controller"
	_ "github.com/erikrios/ponorogo-regency-api/docs"
	"github.com/erikrios/ponorogo-regency-api/middleware"
	"github.com/erikrios/ponorogo-regency-api/repository"
	"github.com/erikrios/ponorogo-regency-api/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           Ponorogo Regency API
// @version         1.0
// @description     API for Administrative Subdivisions of Ponorogo Regency (Districts and Villages).
// @termsOfService  http://swagger.io/terms/

// @contact.name   Erik Rio Setiawan
// @contact.url    http://www.swagger.io/support
// @contact.email  erikriosetiawan15@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host       ponorogo-api.herokuapp.com
// @BasePath   /api/v1
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

	provinceRepository := repository.NewProvinceRepositoryImpl(db)
	regencyRepository := repository.NewRegencyRepositoryImpl(db)
	districtRepository := repository.NewDistrictRepositoryImpl(db)
	villageRepository := repository.NewVillageRepositoryImpl(db)

	provinceService := service.NewProvinceServiceImpl(provinceRepository)
	regencyService := service.NewRegencyServiceImpl(regencyRepository)
	districtService := service.NewDistrictServiceImpl(districtRepository, villageRepository)
	villageService := service.NewVillageServiceImpl(villageRepository)

	provincesController := controller.NewProvincesController(provinceService)
	regenciesController := controller.NewRegenciesController(regencyService)
	districtsController := controller.NewDistrictsController(districtService)
	villagesController := controller.NewVillagesController(villageService)

	e := echo.New()

	if os.Getenv("ENV") == "production" {
		middleware.BodyLimit(e)
		middleware.Gzip(e)
		middleware.RateLimiter(e)
		middleware.Recover(e)
		middleware.Secure(e)
		middleware.RemoveTrailingSlash(e)
	} else {
		middleware.Logger(e)
		middleware.RemoveTrailingSlash(e)
	}

	e.GET("/*", echoSwagger.WrapHandler)

	g := e.Group("/api/v1")
	provincesController.Route(g)
	regenciesController.Route(g)
	districtsController.Route(g)
	villagesController.Route(g)

	e.Logger.Fatal(e.Start(port))
}
