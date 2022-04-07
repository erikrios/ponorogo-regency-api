package service

import (
	"context"

	"github.com/erikrios/ponorogo-regency-api/entity"
	"github.com/erikrios/ponorogo-regency-api/model"
	"github.com/erikrios/ponorogo-regency-api/repository"
)

type villageServiceImpl struct {
	repository repository.VillageRepository
}

func NewVillageServiceImpl(repository repository.VillageRepository) *villageServiceImpl {
	return &villageServiceImpl{repository: repository}
}

func (v *villageServiceImpl) GetAll(ctx context.Context, keyword string) (responses []model.Village, err error) {
	var villages []entity.Village
	var repoErr error

	if keyword == "" {
		villages, repoErr = v.repository.FindAll(ctx)
	} else {
		villages, repoErr = v.repository.FindByName(ctx, keyword)
	}

	if repoErr != nil {
		err = mapError(repoErr)
		return
	}

	responses = make([]model.Village, len(villages))

	for i, village := range villages {
		responses[i] = v.mapToModel(village)
	}
	return
}

func (v *villageServiceImpl) GetByID(ctx context.Context, id string) (response model.Village, err error) {
	village, repoErr := v.repository.FindByID(ctx, id)
	if repoErr != nil {
		err = mapError(repoErr)
		return
	}

	response = v.mapToModel(village)
	return
}

func (v *villageServiceImpl) mapToModel(e entity.Village) model.Village {
	return model.Village{
		ID:   e.ID,
		Name: e.Name,
		District: model.District{
			ID:   e.District.ID,
			Name: e.District.Name,
			Regency: model.Regency{
				ID:   e.District.Regency.ID,
				Name: e.District.Regency.Name,
				Province: model.Province{
					ID:   e.District.Regency.Province.ID,
					Name: e.District.Regency.Province.Name,
				},
			},
		},
	}
}
