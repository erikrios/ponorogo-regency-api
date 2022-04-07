package service

import (
	"context"

	"github.com/erikrios/ponorogo-regency-api/model"
)

type RegencyService interface {
	GetAll(ctx context.Context, keyword string) (responses []model.Regency, err error)
	GetByID(ctx context.Context, id string) (response model.Regency, err error)
}
