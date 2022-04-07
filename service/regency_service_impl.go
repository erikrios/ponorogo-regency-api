package service

import (
	"context"

	"github.com/erikrios/ponorogo-regency-api/entity"
	"github.com/erikrios/ponorogo-regency-api/model"
	"github.com/erikrios/ponorogo-regency-api/repository"
)

type regencyServiceImpl struct {
	repository repository.RegencyRepository
}

func NewRegencyServiceImpl(repository repository.RegencyRepository) *regencyServiceImpl {
	return &regencyServiceImpl{repository: repository}
}

func (r *regencyServiceImpl) GetAll(ctx context.Context, keyword string) (responses []model.Regency, err error) {
	var regencies []entity.Regency
	var repoErr error

	if keyword == "" {
		regencies, repoErr = r.repository.FindAll(ctx)
	} else {
		regencies, repoErr = r.repository.FindByName(ctx, keyword)
	}

	if repoErr != nil {
		err = mapError(repoErr)
		return
	}

	responses = make([]model.Regency, len(regencies))

	for i, regency := range regencies {
		responses[i] = r.mapToModel(regency)
	}
	return
}
func (r *regencyServiceImpl) GetByID(ctx context.Context, id string) (response model.Regency, err error) {
	regency, repoErr := r.repository.FindByID(ctx, id)
	if repoErr != nil {
		err = mapError(repoErr)
		return
	}

	response = r.mapToModel(regency)
	return
}

func (p *regencyServiceImpl) mapToModel(e entity.Regency) model.Regency {
	return model.Regency{
		ID:   e.ID,
		Name: e.Name,
		Province: model.Province{
			ID:   e.Province.ID,
			Name: e.Province.Name,
		},
	}
}
