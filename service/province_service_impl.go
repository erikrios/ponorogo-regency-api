package service

import (
	"context"

	"github.com/erikrios/ponorogo-regency-api/entity"
	"github.com/erikrios/ponorogo-regency-api/model"
	"github.com/erikrios/ponorogo-regency-api/repository"
)

type provinceServiceImpl struct {
	repository repository.ProvinceRepository
}

func NewProvinceServiceImpl(repository repository.ProvinceRepository) *provinceServiceImpl {
	return &provinceServiceImpl{repository: repository}
}

func (p *provinceServiceImpl) GetAll(ctx context.Context, keyword string) (responses []model.Province, err error) {
	var provinces []entity.Province
	var repoErr error

	if keyword == "" {
		provinces, repoErr = p.repository.FindAll(ctx)
	} else {
		provinces, repoErr = p.repository.FindByName(ctx, keyword)
	}

	if repoErr != nil {
		err = mapError(repoErr)
		return
	}

	responses = make([]model.Province, len(provinces))

	for i, province := range provinces {
		responses[i] = p.mapToModel(province)
	}
	return
}

func (p *provinceServiceImpl) GetByID(ctx context.Context, id string) (response model.Province, err error) {
	province, repoErr := p.repository.FindByID(ctx, id)
	if repoErr != nil {
		err = mapError(repoErr)
		return
	}

	response = p.mapToModel(province)
	return
}

func (p *provinceServiceImpl) mapToModel(e entity.Province) model.Province {
	return model.Province{
		ID:   e.ID,
		Name: e.Name,
	}
}
