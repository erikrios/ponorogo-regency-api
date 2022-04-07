package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/erikrios/ponorogo-regency-api/entity"
)

type regencyRepositoryImpl struct {
	db *sql.DB
}

func NewRegencyRepositoryImpl(db *sql.DB) *regencyRepositoryImpl {
	return &regencyRepositoryImpl{db: db}
}

func (r *regencyRepositoryImpl) FindAll(ctx context.Context) (regencies []entity.Regency, err error) {
	statement := "SELECT r.id, r.name, r.province_id, p.name AS province_name FROM regencies r INNER JOIN provinces p on r.province_id = p.id;"

	rows, err := r.db.QueryContext(ctx, statement)
	if err != nil {
		log.Println(err)
		err = ErrDatabase
		return
	}

	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			log.Println(err.Error())
		}
	}(rows)

	regencies = make([]entity.Regency, 0)
	for rows.Next() {
		var regency entity.Regency
		if err = rows.Scan(&regency.ID, &regency.Name, &regency.Province.ID, &regency.Province.Name); err != nil {
			log.Println(err)
			err = ErrDatabase
			return
		}
		regencies = append(regencies, regency)
	}

	return
}

func (r *regencyRepositoryImpl) FindByID(ctx context.Context, id string) (regency entity.Regency, err error) {
	statement := "SELECT r.id, r.name, r.province_id, p.name AS province_name FROM regencies r INNER JOIN provinces p on r.province_id = p.id WHERE r.id = $1;"

	row := r.db.QueryRowContext(ctx, statement, id)

	switch scanErr := row.Scan(&regency.ID, &regency.Name, &regency.Province.ID, &regency.Province.Name); scanErr {
	case sql.ErrNoRows:
		err = ErrQueryNotFound
		return
	case nil:
		return
	default:
		err = ErrDatabase
		log.Println(scanErr)
		return
	}
}

func (r *regencyRepositoryImpl) FindByName(ctx context.Context, keyword string) (regencies []entity.Regency, err error) {
	statement := "SELECT r.id, r.name, r.province_id, p.name AS province_name FROM regencies r INNER JOIN provinces p on r.province_id = p.id WHERE r.name ILIKE '%$1%';"

	rows, err := r.db.QueryContext(ctx, statement, keyword)
	if err != nil {
		log.Println(err)
		err = ErrDatabase
		return
	}

	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			log.Println(err.Error())
		}
	}(rows)

	regencies = make([]entity.Regency, 0)
	for rows.Next() {
		var regency entity.Regency
		if err = rows.Scan(&regency.ID, &regency.Name, &regency.Province.ID, &regency.Province.Name); err != nil {
			log.Println(err)
			err = ErrDatabase
			return
		}
		regencies = append(regencies, regency)
	}

	return
}
