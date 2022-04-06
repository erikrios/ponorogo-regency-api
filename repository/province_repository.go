package repository

import (
	"context"

	"github.com/erikrios/ponorogo-regency-api/entity"
)

type ProvinceRepository interface {
	FindAll(ctx context.Context) (provinces []entity.Province, err error)
	FindByID(ctx context.Context, id string) (province entity.Province, err error)
	FindByName(ctx context.Context, keyword string) (provinces []entity.Province, err error)
}
