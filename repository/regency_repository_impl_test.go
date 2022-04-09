package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/erikrios/ponorogo-regency-api/entity"
	"github.com/stretchr/testify/assert"
)

func TestRegencyRepositoryImpl(t *testing.T) {

	t.Run("TestFindAll", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		expectedRegencies := []entity.Regency{
			{
				ID:   "3201",
				Name: "Kabupaten Ponorogo",
				Province: entity.Province{
					ID:   "32",
					Name: "Jawa Timur",
				},
			},
		}

		returnedRows := sqlmock.NewRows([]string{"id", "name", "province_id", "province_name"})
		for _, regency := range expectedRegencies {
			returnedRows.AddRow(regency.ID, regency.Name, regency.Province.ID, regency.Province.Name)
		}

		t.Run("it should return valid regencies, when database successfully return the data", func(t *testing.T) {
			mock.ExpectQuery("SELECT r.id, r.name, r.province_id, p.name AS province_name FROM regencies r INNER JOIN provinces p on r.province_id = p.id;").WillReturnRows(returnedRows)

			var repo RegencyRepository = NewRegencyRepositoryImpl(db)

			got, err := repo.FindAll(context.Background())
			if err != nil {
				t.Fatal(err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}

			assert.ElementsMatch(t, expectedRegencies, got)
		})

		t.Run("it should return error, when database return an error", func(t *testing.T) {
			mock.ExpectQuery("SELECT r.id, r.name, r.province_id, p.name AS province_name FROM regencies r INNER JOIN provinces p on r.province_id = p.id;").WillReturnError(ErrDatabase)

			var repo RegencyRepository = NewRegencyRepositoryImpl(db)

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

		expectedRegency := entity.Regency{
			ID:   "3201",
			Name: "Kabupaten Ponorogo",
			Province: entity.Province{
				ID:   "32",
				Name: "Jawa Timur",
			},
		}

		returnedRows := sqlmock.NewRows([]string{"id", "name", "regency_id", "regency_name"})
		returnedRows.AddRow(expectedRegency.ID, expectedRegency.Name, expectedRegency.Province.ID, expectedRegency.Province.Name)

		t.Run("it should return valid regency, when database successfully return the data", func(t *testing.T) {
			mock.ExpectQuery(`SELECT r.id, r.name, r.province_id, p.name AS province_name FROM regencies r INNER JOIN provinces p on r.province_id = p.id WHERE r.id = \$1;`).WithArgs(expectedRegency.ID).WillReturnRows(returnedRows)

			var repo RegencyRepository = NewRegencyRepositoryImpl(db)

			got, err := repo.FindByID(context.Background(), expectedRegency.ID)
			if err != nil {
				t.Fatal(err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, expectedRegency, got)
		})

		t.Run("it should return error, when database return an error", func(t *testing.T) {
			mock.ExpectQuery(`SELECT r.id, r.name, r.province_id, p.name AS province_name FROM regencies r INNER JOIN provinces p on r.province_id = p.id WHERE r.id = \$1;`).WithArgs(expectedRegency.ID).WillReturnError(ErrDatabase)

			var repo RegencyRepository = NewRegencyRepositoryImpl(db)

			if _, err := repo.FindByID(context.Background(), expectedRegency.ID); assert.Error(t, err) {
				assert.Equal(t, ErrDatabase, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}
		})

		t.Run("it should return not found error, when given id not found in the  database", func(t *testing.T) {
			mock.ExpectQuery(`SELECT r.id, r.name, r.province_id, p.name AS province_name FROM regencies r INNER JOIN provinces p on r.province_id = p.id WHERE r.id = \$1;`).WithArgs(expectedRegency.ID).WillReturnError(sql.ErrNoRows)

			var repo RegencyRepository = NewRegencyRepositoryImpl(db)

			if _, err := repo.FindByID(context.Background(), expectedRegency.ID); assert.Error(t, err) {
				assert.Equal(t, ErrQueryNotFound, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}
		})
	})

	t.Run("TestFindByName", func(t *testing.T) {})

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	expectedRegencies := []entity.Regency{
		{
			ID:   "3201",
			Name: "Kabupaten Ponorogo",
			Province: entity.Province{
				ID:   "32",
				Name: "Jawa Timur",
			},
		},
	}

	returnedRows := sqlmock.NewRows([]string{"id", "name", "province_id", "province_name"})
	for _, regency := range expectedRegencies {
		returnedRows.AddRow(regency.ID, regency.Name, regency.Province.ID, regency.Province.Name)
	}

	t.Run("it should return valid regencies, when database successfully return the data", func(t *testing.T) {
		mock.ExpectQuery(`SELECT r.id, r.name, r.province_id, p.name AS province_name FROM regencies r INNER JOIN provinces p on r.province_id = p.id WHERE r.name ILIKE '%' || \$1 || '%';`).WithArgs(expectedRegencies[0].Name).WillReturnRows(returnedRows)

		var repo RegencyRepository = NewRegencyRepositoryImpl(db)

		got, err := repo.FindByName(context.Background(), expectedRegencies[0].Name)
		if err != nil {
			t.Fatal(err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Fatal(err)
		}

		assert.ElementsMatch(t, expectedRegencies, got)
	})

	t.Run("it should return error, when database return an error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT r.id, r.name, r.province_id, p.name AS province_name FROM regencies r INNER JOIN provinces p on r.province_id = p.id WHERE r.name ILIKE '%' || \$1 || '%';`).WithArgs(expectedRegencies[0].Name).WillReturnError(ErrDatabase)

		var repo RegencyRepository = NewRegencyRepositoryImpl(db)

		if _, err := repo.FindByName(context.Background(), expectedRegencies[0].Name); assert.Error(t, err) {
			assert.Equal(t, ErrDatabase, err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Fatal(err)
		}
	})
}
