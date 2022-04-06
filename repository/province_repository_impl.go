package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/erikrios/ponorogo-regency-api/entity"
)

type provinceRepositoryImpl struct {
	db *sql.DB
}

func NewProvinceRepositoryImpl(db *sql.DB) *provinceRepositoryImpl {
	return &provinceRepositoryImpl{db: db}
}

func (p *provinceRepositoryImpl) FindAll(ctx context.Context) (provinces []entity.Province, err error) {
	statement := "SELECT p.id, p.name FROM provinces p;"

	rows, err := p.db.QueryContext(ctx, statement)
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

	provinces = make([]entity.Province, 0)
	for rows.Next() {
		var province entity.Province
		if err = rows.Scan(&province.ID, &province.Name); err != nil {
			log.Println(err)
			err = ErrDatabase
			return
		}
		provinces = append(provinces, province)
	}

	return
}

func (p *provinceRepositoryImpl) FindByID(ctx context.Context, id string) (province entity.Province, err error) {
	return
}

func (p *provinceRepositoryImpl) FindByName(ctx context.Context, keyword string) (provinces []entity.Province, err error) {
	return
}
