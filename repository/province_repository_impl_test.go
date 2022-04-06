package repository

import (
	"context"
	"log"
	"testing"

	"github.com/erikrios/ponorogo-regency-api/config"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./../.env"); err != nil {
		log.Fatalf("Error loading .env file: %s\n", err.Error())
	}
}

func TestFindAll(t *testing.T) {
	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		t.Fatal(err)
	}
	var repo ProvinceRepository = NewProvinceRepositoryImpl(db)

	provinces, err := repo.FindAll(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Log(provinces)
}
