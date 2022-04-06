package repository

import (
	"context"

	"github.com/erikrios/ponorogo-regency-api/entity"
)

type RegencyRepository interface {
	FindAll(ctx context.Context) (regencies []entity.Regency, err error)
	FindByID(ctx context.Context, id string) (regency entity.Regency, err error)
	FindByName(ctx context.Context, keyword string) (regencies []entity.Regency, err error)
}
