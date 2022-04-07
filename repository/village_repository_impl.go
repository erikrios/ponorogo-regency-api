package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/erikrios/ponorogo-regency-api/entity"
)

type villageRepositoryImpl struct {
	db *sql.DB
}

func NewVillageRepositoryImpl(db *sql.DB) *villageRepositoryImpl {
	return &villageRepositoryImpl{db: db}
}

func (v *villageRepositoryImpl) FindAll(ctx context.Context) (villages []entity.Village, err error) {
	statement := "SELECT v.id, v.name, v.district_id, d.name AS district_name, d.regency_id, r.name AS regency_name, r.province_id, p.name AS province_name FROM villages v INNER JOIN districts d on d.id = v.district_id INNER JOIN regencies r on d.regency_id = r.id INNER JOIN provinces p on r.province_id = p.id;"

	rows, err := v.db.QueryContext(ctx, statement)
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

	villages = make([]entity.Village, 0)
	for rows.Next() {
		var village entity.Village
		if err = rows.Scan(
			&village.ID,
			&village.Name,
			&village.District.ID,
			&village.District.Name,
			&village.District.Regency.ID,
			&village.District.Regency.Name,
			&village.District.Regency.Province.ID,
			&village.District.Regency.Province.Name,
		); err != nil {
			log.Println(err)
			err = ErrDatabase
			return
		}
		villages = append(villages, village)
	}

	return
}

func (v *villageRepositoryImpl) FindByID(ctx context.Context, id string) (village entity.Village, err error) {
	statement := "SELECT v.id, v.name, v.district_id, d.name AS district_name, d.regency_id, r.name AS regency_name, r.province_id, p.name AS province_name FROM villages v INNER JOIN districts d on d.id = v.district_id INNER JOIN regencies r on d.regency_id = r.id INNER JOIN provinces p on r.province_id = p.id WHERE v.id = $1;"

	row := v.db.QueryRowContext(ctx, statement, id)

	switch scanErr := row.Scan(
		&village.ID,
		&village.Name,
		&village.District.ID,
		&village.District.Name,
		&village.District.Regency.ID,
		&village.District.Regency.Name,
		&village.District.Regency.Province.ID,
		&village.District.Regency.Province.Name,
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

func (v *villageRepositoryImpl) FindByName(ctx context.Context, keyword string) (villages []entity.Village, err error) {
	statement := "SELECT v.id, v.name, v.district_id, d.name AS district_name, d.regency_id, r.name AS regency_name, r.province_id, p.name AS province_name FROM villages v INNER JOIN districts d on d.id = v.district_id INNER JOIN regencies r on d.regency_id = r.id INNER JOIN provinces p on r.province_id = p.id WHERE v.name ILIKE '%' || $1 || '%';"

	rows, err := v.db.QueryContext(ctx, statement, keyword)
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

	villages = make([]entity.Village, 0)
	for rows.Next() {
		var village entity.Village
		if err = rows.Scan(
			&village.ID,
			&village.Name,
			&village.District.ID,
			&village.District.Name,
			&village.District.Regency.ID,
			&village.District.Regency.Name,
			&village.District.Regency.Province.ID,
			&village.District.Regency.Province.Name,
		); err != nil {
			log.Println(err)
			err = ErrDatabase
			return
		}
		villages = append(villages, village)
	}

	return
}

func (v *villageRepositoryImpl) FindByDistrictID(ctx context.Context, districtID string) (villages []entity.Village, err error) {
	statement := "SELECT v.id, v.name, v.district_id, d.name AS district_name, d.regency_id, r.name AS regency_name, r.province_id, p.name AS province_name FROM villages v INNER JOIN districts d on d.id = v.district_id INNER JOIN regencies r on d.regency_id = r.id INNER JOIN provinces p on r.province_id = p.id WHERE v.district_id = $1;"

	rows, err := v.db.QueryContext(ctx, statement, districtID)
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

	villages = make([]entity.Village, 0)
	for rows.Next() {
		var village entity.Village
		if err = rows.Scan(
			&village.ID,
			&village.Name,
			&village.District.ID,
			&village.District.Name,
			&village.District.Regency.ID,
			&village.District.Regency.Name,
			&village.District.Regency.Province.ID,
			&village.District.Regency.Province.Name,
		); err != nil {
			log.Println(err)
			err = ErrDatabase
			return
		}
		villages = append(villages, village)
	}

	return
}

func (v *villageRepositoryImpl) FindByDistrictName(ctx context.Context, keyword string) (villages []entity.Village, err error) {
	statement := "SELECT v.id, v.name, v.district_id, d.name AS district_name, d.regency_id, r.name AS regency_name, r.province_id, p.name AS province_name FROM villages v INNER JOIN districts d on d.id = v.district_id INNER JOIN regencies r on d.regency_id = r.id INNER JOIN provinces p on r.province_id = p.id WHERE d.name ILIKE '%' || $1 || '%';"

	rows, err := v.db.QueryContext(ctx, statement, keyword)
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

	villages = make([]entity.Village, 0)
	for rows.Next() {
		var village entity.Village
		if err = rows.Scan(
			&village.ID,
			&village.Name,
			&village.District.ID,
			&village.District.Name,
			&village.District.Regency.ID,
			&village.District.Regency.Name,
			&village.District.Regency.Province.ID,
			&village.District.Regency.Province.Name,
		); err != nil {
			log.Println(err)
			err = ErrDatabase
			return
		}
		villages = append(villages, village)
	}

	return
}
