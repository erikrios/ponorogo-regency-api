package service

import (
	"context"

	"github.com/erikrios/ponorogo-regency-api/model"
)

type VillageService interface {
	GetAll(ctx context.Context, keyword string) (responses []model.Village, err error)
	GetByID(ctx context.Context, id string) (response model.Village, err error)
}
