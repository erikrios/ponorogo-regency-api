package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/erikrios/ponorogo-regency-api/entity"
)

type districtRepositoryImpl struct {
	db *sql.DB
}

func NewDistrictRepositoryImpl(db *sql.DB) *districtRepositoryImpl {
	return &districtRepositoryImpl{db: db}
}

func (d *districtRepositoryImpl) FindAll(ctx context.Context) (districts []entity.District, err error) {
	statement := "SELECT d.id, d.name, d.regency_id, r.name AS regency_name, r.province_id, p.name AS province_name FROM districts d INNER JOIN regencies r on d.regency_id = r.id INNER JOIN provinces p on r.province_id = p.id;"

	rows, err := d.db.QueryContext(ctx, statement)
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

	districts = make([]entity.District, 0)
	for rows.Next() {
		var district entity.District
		if err = rows.Scan(
			&district.ID,
			&district.Name,
			&district.Regency.ID,
			&district.Regency.Name,
			&district.Regency.Province.ID,
			&district.Regency.Province.Name,
		); err != nil {
			log.Println(err)
			err = ErrDatabase
			return
		}
		districts = append(districts, district)
	}

	return
}

func (d *districtRepositoryImpl) FindByID(ctx context.Context, id string) (district entity.District, err error) {
	statement := "SELECT d.id, d.name, d.regency_id, r.name AS regency_name, r.province_id, p.name AS province_name FROM districts d INNER JOIN regencies r on d.regency_id = r.id INNER JOIN provinces p on r.province_id = p.id WHERE d.id = $1;"

	row := d.db.QueryRowContext(ctx, statement, id)

	switch scanErr := row.Scan(
		&district.ID,
		&district.Name,
		&district.Regency.ID,
		&district.Regency.Name,
		&district.Regency.Province.ID,
		&district.Regency.Province.Name,
	); scanErr {
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

func (d *districtRepositoryImpl) FindByName(ctx context.Context, keyword string) (districts []entity.District, err error) {
	statement := "SELECT d.id, d.name, d.regency_id, r.name AS regency_name, r.province_id, p.name AS province_name FROM districts d INNER JOIN regencies r on d.regency_id = r.id INNER JOIN provinces p on r.province_id = p.id WHERE d.name ILIKE '%' || $1 || '%';"

	rows, err := d.db.QueryContext(ctx, statement, keyword)
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

	districts = make([]entity.District, 0)
	for rows.Next() {
		var district entity.District
		if err = rows.Scan(
			&district.ID,
			&district.Name,
			&district.Regency.ID,
			&district.Regency.Name,
			&district.Regency.Province.ID,
			&district.Regency.Province.Name,
		); err != nil {
			log.Println(err)
			err = ErrDatabase
			return
		}
		districts = append(districts, district)
	}

	return
}
