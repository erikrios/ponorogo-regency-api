package repository

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/erikrios/ponorogo-regency-api/entity"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	if err := godotenv.Load("./../.env.local"); err != nil {
		log.Fatalf("Error loading .env file: %s\n", err.Error())
	}
}

func TestFindAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	expectedProvinces := []entity.Province{
		{
			ID:   "32",
			Name: "Jawa Timur",
		},
	}

	returnedRows := sqlmock.NewRows([]string{"id", "name"})
	for _, province := range expectedProvinces {
		returnedRows.AddRow(province.ID, province.Name)
	}

	t.Run("it should return valid provinces, when database successfully return the data", func(t *testing.T) {
		mock.ExpectQuery("SELECT p.id, p.name FROM provinces p;").WillReturnRows(returnedRows)

		var repo ProvinceRepository = NewProvinceRepositoryImpl(db)

		got, err := repo.FindAll(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Fatal(err)
		}

		assert.ElementsMatch(t, expectedProvinces, got)
	})

	t.Run("it should return error, when database return an error", func(t *testing.T) {
		mock.ExpectQuery("SELECT p.id, p.name FROM provinces p;").WillReturnError(ErrDatabase)

		var repo ProvinceRepository = NewProvinceRepositoryImpl(db)

		if _, err := repo.FindAll(context.Background()); assert.Error(t, err) {
			assert.Equal(t, ErrDatabase, err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Fatal(err)
		}
	})
}

func TestFindByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	expectedProvince := entity.Province{
		ID:   "32",
		Name: "Jawa Timur",
	}

	returnedRows := sqlmock.NewRows([]string{"id", "name"})
	returnedRows.AddRow(expectedProvince.ID, expectedProvince.Name)

	t.Run("it should return valid province, when database successfully return the data", func(t *testing.T) {
		mock.ExpectQuery(`SELECT p.id, p.name FROM provinces p WHERE p.id = \$1;`).WithArgs(expectedProvince.ID).WillReturnRows(returnedRows)

		var repo ProvinceRepository = NewProvinceRepositoryImpl(db)

		got, err := repo.FindByID(context.Background(), expectedProvince.ID)
		if err != nil {
			t.Fatal(err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, expectedProvince, got)
	})

	t.Run("it should return error, when database return an error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT p.id, p.name FROM provinces p WHERE p.id = \$1;`).WithArgs(expectedProvince.ID).WillReturnError(ErrDatabase)

		var repo ProvinceRepository = NewProvinceRepositoryImpl(db)

		if _, err := repo.FindByID(context.Background(), expectedProvince.ID); assert.Error(t, err) {
			assert.Equal(t, ErrDatabase, err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("it should return not found error, when given id not found in the  database", func(t *testing.T) {
		mock.ExpectQuery(`SELECT p.id, p.name FROM provinces p WHERE p.id = \$1;`).WithArgs(expectedProvince.ID).WillReturnError(sql.ErrNoRows)

		var repo ProvinceRepository = NewProvinceRepositoryImpl(db)

		if _, err := repo.FindByID(context.Background(), expectedProvince.ID); assert.Error(t, err) {
			assert.Equal(t, ErrQueryNotFound, err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Fatal(err)
		}
	})
}

func TestFindByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	expectedProvinces := []entity.Province{
		{
			ID:   "32",
			Name: "Jawa Timur",
		},
	}

	returnedRows := sqlmock.NewRows([]string{"id", "name"})
	for _, province := range expectedProvinces {
		returnedRows.AddRow(province.ID, province.Name)
	}
	t.Run("it should return valid provinces, when database successfully return the data", func(t *testing.T) {
		mock.ExpectQuery(`SELECT p.id, p.name FROM provinces p WHERE p.name ILIKE \$1;`).WithArgs(expectedProvinces[0].Name).WillReturnRows(returnedRows)

		var repo ProvinceRepository = NewProvinceRepositoryImpl(db)

		got, err := repo.FindByName(context.Background(), expectedProvinces[0].Name)
		if err != nil {
			t.Fatal(err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Fatal(err)
		}

		assert.ElementsMatch(t, expectedProvinces, got)
	})

	t.Run("it should return error, when database return an error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT p.id, p.name FROM provinces p WHERE p.name ILIKE \$1;`).WithArgs(expectedProvinces[0].Name).WillReturnError(ErrDatabase)

		var repo ProvinceRepository = NewProvinceRepositoryImpl(db)

		if _, err := repo.FindByName(context.Background(), expectedProvinces[0].Name); assert.Error(t, err) {
			assert.Equal(t, ErrDatabase, err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Fatal(err)
		}
	})
}
