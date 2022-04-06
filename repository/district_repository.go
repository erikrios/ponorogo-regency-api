package repository

import (
	"context"

	"github.com/erikrios/ponorogo-regency-api/entity"
)

type DistrictRepository interface {
	FindAll(ctx context.Context) (districts []entity.District, err error)
	FindByID(ctx context.Context, id string) (district entity.District, err error)
	FindByName(ctx context.Context, keyword string) (districts []entity.District, err error)
}
