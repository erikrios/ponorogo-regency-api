package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/erikrios/ponorogo-regency-api/entity"
	"github.com/stretchr/testify/assert"
)

func TestVillageRepositoryImpl(t *testing.T) {

	t.Run("TestFindAll", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		expectedVillages := []entity.Village{
			{
				ID:   "4050101101",
				Name: "Pager",
				District: entity.District{
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
			},
		}

		returnedRows := sqlmock.NewRows([]string{"id", "name", "district_id", "district_name", "regency_id", "regency_name", "province_id", "province_name"})
		for _, village := range expectedVillages {
			returnedRows.AddRow(
				village.ID,
				village.Name,
				village.District.ID,
				village.District.Name,
				village.District.Regency.ID,
				village.District.Regency.Name,
				village.District.Regency.Province.ID,
				village.District.Regency.Province.Name,
			)
		}

		t.Run("it should return valid villages, when database successfully return the data", func(t *testing.T) {
			mock.ExpectQuery(".*").WillReturnRows(returnedRows)

			var repo VillageRepository = NewVillageRepositoryImpl(db)

			got, err := repo.FindAll(context.Background())
			if err != nil {
				t.Fatal(err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}

			assert.ElementsMatch(t, expectedVillages, got)
		})

		t.Run("it should return error, when database return an error", func(t *testing.T) {
			mock.ExpectQuery(".*").WillReturnError(ErrDatabase)

			var repo VillageRepository = NewVillageRepositoryImpl(db)

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

		expectedVillage := entity.Village{
			ID:   "4050101101",
			Name: "Pager",
			District: entity.District{
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

		returnedRows := sqlmock.NewRows([]string{"id", "name", "district_id", "district_name", "regency_id", "regency_name", "province_id", "province_name"})
		returnedRows.AddRow(
			expectedVillage.ID,
			expectedVillage.Name,
			expectedVillage.District.ID,
			expectedVillage.District.Name,
			expectedVillage.District.Regency.ID,
			expectedVillage.District.Regency.Name,
			expectedVillage.District.Regency.Province.ID,
			expectedVillage.District.Regency.Province.Name,
		)

		t.Run("it should return valid village, when database successfully return the data", func(t *testing.T) {
			mock.ExpectQuery(".*").WithArgs(expectedVillage.ID).WillReturnRows(returnedRows)

			var repo VillageRepository = NewVillageRepositoryImpl(db)

			got, err := repo.FindByID(context.Background(), expectedVillage.ID)
			if err != nil {
				t.Fatal(err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, expectedVillage, got)
		})

		t.Run("it should return error, when database return an error", func(t *testing.T) {
			mock.ExpectQuery(".*").WithArgs(expectedVillage.ID).WillReturnError(ErrDatabase)

			var repo VillageRepository = NewVillageRepositoryImpl(db)

			if _, err := repo.FindByID(context.Background(), expectedVillage.ID); assert.Error(t, err) {
				assert.Equal(t, ErrDatabase, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}
		})

		t.Run("it should return not found error, when given id not found in the  database", func(t *testing.T) {
			mock.ExpectQuery(".*").WithArgs(expectedVillage.ID).WillReturnError(sql.ErrNoRows)

			var repo VillageRepository = NewVillageRepositoryImpl(db)

			if _, err := repo.FindByID(context.Background(), expectedVillage.ID); assert.Error(t, err) {
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

		expectedVillages := []entity.Village{
			{
				ID:   "4050101101",
				Name: "Pager",
				District: entity.District{
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
			},
		}

		returnedRows := sqlmock.NewRows([]string{"id", "name", "district_id", "district_name", "regency_id", "regency_name", "province_id", "province_name"})
		for _, village := range expectedVillages {
			returnedRows.AddRow(
				village.ID,
				village.Name,
				village.District.ID,
				village.District.Name,
				village.District.Regency.ID,
				village.District.Regency.Name,
				village.District.Regency.Province.ID,
				village.District.Regency.Province.Name,
			)
		}

		t.Run("it should return valid villages, when database successfully return the data", func(t *testing.T) {
			mock.ExpectQuery(".*").WithArgs(expectedVillages[0].Name).WillReturnRows(returnedRows)

			var repo VillageRepository = NewVillageRepositoryImpl(db)

			got, err := repo.FindByName(context.Background(), expectedVillages[0].Name)
			if err != nil {
				t.Fatal(err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}

			assert.ElementsMatch(t, expectedVillages, got)
		})

		t.Run("it should return error, when database return an error", func(t *testing.T) {
			mock.ExpectQuery(".*").WithArgs(expectedVillages[0].Name).WillReturnError(ErrDatabase)

			var repo VillageRepository = NewVillageRepositoryImpl(db)

			if _, err := repo.FindByName(context.Background(), expectedVillages[0].Name); assert.Error(t, err) {
				assert.Equal(t, ErrDatabase, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}
		})
	})

	t.Run("TestFindByDistrictID", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		expectedVillages := []entity.Village{
			{
				ID:   "4050101101",
				Name: "Pager",
				District: entity.District{
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
			},
		}

		returnedRows := sqlmock.NewRows([]string{"id", "name", "district_id", "district_name", "regency_id", "regency_name", "province_id", "province_name"})
		for _, village := range expectedVillages {
			returnedRows.AddRow(
				village.ID,
				village.Name,
				village.District.ID,
				village.District.Name,
				village.District.Regency.ID,
				village.District.Regency.Name,
				village.District.Regency.Province.ID,
				village.District.Regency.Province.Name,
			)
		}

		t.Run("it should return valid villages, when database successfully return the data", func(t *testing.T) {
			mock.ExpectQuery(".*").WithArgs(expectedVillages[0].District.ID).WillReturnRows(returnedRows)

			var repo VillageRepository = NewVillageRepositoryImpl(db)

			got, err := repo.FindByDistrictID(context.Background(), expectedVillages[0].District.ID)
			if err != nil {
				t.Fatal(err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}

			assert.ElementsMatch(t, expectedVillages, got)
		})

		t.Run("it should return error, when database return an error", func(t *testing.T) {
			mock.ExpectQuery(".*").WithArgs(expectedVillages[0].District.ID).WillReturnError(ErrDatabase)

			var repo VillageRepository = NewVillageRepositoryImpl(db)

			if _, err := repo.FindByDistrictID(context.Background(), expectedVillages[0].District.ID); assert.Error(t, err) {
				assert.Equal(t, ErrDatabase, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}
		})
	})

	t.Run("TestFindByDistrictName", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		expectedVillages := []entity.Village{
			{
				ID:   "4050101101",
				Name: "Pager",
				District: entity.District{
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
			},
		}

		returnedRows := sqlmock.NewRows([]string{"id", "name", "district_id", "district_name", "regency_id", "regency_name", "province_id", "province_name"})
		for _, village := range expectedVillages {
			returnedRows.AddRow(
				village.ID,
				village.Name,
				village.District.ID,
				village.District.Name,
				village.District.Regency.ID,
				village.District.Regency.Name,
				village.District.Regency.Province.ID,
				village.District.Regency.Province.Name,
			)
		}

		t.Run("it should return valid villages, when database successfully return the data", func(t *testing.T) {
			mock.ExpectQuery(".*").WithArgs(expectedVillages[0].District.Name).WillReturnRows(returnedRows)

			var repo VillageRepository = NewVillageRepositoryImpl(db)

			got, err := repo.FindByDistrictName(context.Background(), expectedVillages[0].District.Name)
			if err != nil {
				t.Fatal(err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}

			assert.ElementsMatch(t, expectedVillages, got)
		})

		t.Run("it should return error, when database return an error", func(t *testing.T) {
			mock.ExpectQuery(".*").WithArgs(expectedVillages[0].District.Name).WillReturnError(ErrDatabase)

			var repo VillageRepository = NewVillageRepositoryImpl(db)

			if _, err := repo.FindByDistrictName(context.Background(), expectedVillages[0].District.Name); assert.Error(t, err) {
				assert.Equal(t, ErrDatabase, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}
		})
	})
}
