package repository

import (
	"context"

	"github.com/erikrios/ponorogo-regency-api/entity"
)

type VillageRepository interface {
	FindAll(ctx context.Context) (villages []entity.Village, err error)
	FindByID(ctx context.Context, id string) (village entity.Village, err error)
	FindByName(ctx context.Context, keyword string) (villages []entity.Village, err error)
	FindByDistrictID(ctx context.Context, districtID string) (villages []entity.Village, err error)
	FindByDistrictName(ctx context.Context, keyword string) (villages []entity.Village, err error)
}
