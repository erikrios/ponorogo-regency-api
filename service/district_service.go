package service

import (
	"context"

	"github.com/erikrios/ponorogo-regency-api/model"
)

type DistrictService interface {
	GetAll(ctx context.Context, keyword string) (responses []model.District, err error)
	GetByID(ctx context.Context, id string) (response model.District, err error)
	GetVillagesByDistrictID(ctx context.Context, id string) (responses []model.Village, err error)
	GetVillagesByDistrictName(ctx context.Context, keyword string) (responses []model.Village, err error)
}
