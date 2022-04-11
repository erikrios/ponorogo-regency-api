package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/erikrios/ponorogo-regency-api/entity"
	"github.com/stretchr/testify/assert"
)

func TestDistrictRepositoryImpl(t *testing.T) {

	t.Run("TestFindAll", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		expectedDistricts := []entity.District{
			{
				ID:   "32010001",
				Name: "Bungkal",
				Regency: entity.Regency{
					ID:   "3201",
					Name: "Kabupaten Ponorogo",
					Province: entity.Province{
						ID:   "32",
						Name: "Jawa Timur",
					},
				},
			},
		}

		returnedRows := sqlmock.NewRows([]string{"id", "name", "regency_id", "regency_name", "province_id", "province_name"})
		for _, district := range expectedDistricts {
			returnedRows.AddRow(district.ID, district.Name, district.Regency.ID, district.Regency.Name, district.Regency.Province.ID, district.Regency.Province.Name)
		}

		t.Run("it should return valid districts, when database successfully return the data", func(t *testing.T) {
			mock.ExpectQuery(".*").WillReturnRows(returnedRows)

			var repo DistrictRepository = NewDistrictRepositoryImpl(db)

			got, err := repo.FindAll(context.Background())
			if err != nil {
				t.Fatal(err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}

			assert.ElementsMatch(t, expectedDistricts, got)
		})

		t.Run("it should return error, when database return an error", func(t *testing.T) {
			mock.ExpectQuery(".*").WillReturnError(ErrDatabase)

			var repo DistrictRepository = NewDistrictRepositoryImpl(db)

			if _, err := repo.FindAll(context.Background()); assert.Error(t, err) {
				assert.Equal(t, ErrDatabase, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}
		})
	})

	t.Run("TestFindByID", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		expectedDistrict := entity.District{
			ID:   "32010001",
			Name: "Bungkal",
			Regency: entity.Regency{
				ID:   "3201",
				Name: "Kabupaten Ponorogo",
				Province: entity.Province{
					ID:   "32",
					Name: "Jawa Timur",
				},
			},
		}

		returnedRows := sqlmock.NewRows([]string{"id", "name", "regency_id", "regency_name", "province_id", "province_name"})
		returnedRows.AddRow(expectedDistrict.ID, expectedDistrict.Name, expectedDistrict.Regency.ID, expectedDistrict.Regency.Name, expectedDistrict.Regency.Province.ID, expectedDistrict.Regency.Province.Name)

		t.Run("it should return valid district, when database successfully return the data", func(t *testing.T) {
			mock.ExpectQuery(".*").WithArgs(expectedDistrict.ID).WillReturnRows(returnedRows)

			var repo DistrictRepository = NewDistrictRepositoryImpl(db)

			got, err := repo.FindByID(context.Background(), expectedDistrict.ID)
			if err != nil {
				t.Fatal(err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, expectedDistrict, got)
		})

		t.Run("it should return error, when database return an error", func(t *testing.T) {
			mock.ExpectQuery(".*").WithArgs(expectedDistrict.ID).WillReturnError(ErrDatabase)

			var repo DistrictRepository = NewDistrictRepositoryImpl(db)

			if _, err := repo.FindByID(context.Background(), expectedDistrict.ID); assert.Error(t, err) {
				assert.Equal(t, ErrDatabase, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}
		})

		t.Run("it should return not found error, when given id not found in the  database", func(t *testing.T) {
			mock.ExpectQuery(".*").WithArgs(expectedDistrict.ID).WillReturnError(sql.ErrNoRows)

			var repo DistrictRepository = NewDistrictRepositoryImpl(db)

			if _, err := repo.FindByID(context.Background(), expectedDistrict.ID); assert.Error(t, err) {
				assert.Equal(t, ErrQueryNotFound, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}
		})
	})

	t.Run("TestFindByName", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		expectedDistricts := []entity.District{
			{
				ID:   "32010001",
				Name: "Bungkal",
				Regency: entity.Regency{
					ID:   "3201",
					Name: "Kabupaten Ponorogo",
					Province: entity.Province{
						ID:   "32",
						Name: "Jawa Timur",
					},
				},
			},
		}

		returnedRows := sqlmock.NewRows([]string{"id", "name", "regency_id", "regency_name", "province_id", "province_name"})
		for _, district := range expectedDistricts {
			returnedRows.AddRow(district.ID, district.Name, district.Regency.ID, district.Regency.Name, district.Regency.Province.ID, district.Regency.Province.Name)
		}

		t.Run("it should return valid districts, when database successfully return the data", func(t *testing.T) {
			mock.ExpectQuery(".*").WithArgs(expectedDistricts[0].Name).WillReturnRows(returnedRows)

			var repo DistrictRepository = NewDistrictRepositoryImpl(db)

			got, err := repo.FindByName(context.Background(), expectedDistricts[0].Name)
			if err != nil {
				t.Fatal(err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}

			assert.ElementsMatch(t, expectedDistricts, got)
		})

		t.Run("it should return error, when database return an error", func(t *testing.T) {
			mock.ExpectQuery(".*").WithArgs(expectedDistricts[0].Name).WillReturnError(ErrDatabase)

			var repo DistrictRepository = NewDistrictRepositoryImpl(db)

			if _, err := repo.FindByName(context.Background(), expectedDistricts[0].Name); assert.Error(t, err) {
				assert.Equal(t, ErrDatabase, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}
		})
	})
}
