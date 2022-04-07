package service

import (
	"context"

	"github.com/erikrios/ponorogo-regency-api/model"
)

type ProvinceService interface {
	GetAll(ctx context.Context, keyword string) (responses []model.Province, err error)
	GetByID(ctx context.Context, id string) (response model.Province, err error)
}
