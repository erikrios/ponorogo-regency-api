package service

import (
	"context"

	"github.com/erikrios/ponorogo-regency-api/entity"
	"github.com/erikrios/ponorogo-regency-api/model"
	"github.com/erikrios/ponorogo-regency-api/repository"
)

type districtServiceImpl struct {
	districtRepository repository.DistrictRepository
	villageRepository  repository.VillageRepository
}

func NewDistrictServiceImpl(
	districtRepository repository.DistrictRepository,
	villageRepository repository.VillageRepository,
) *districtServiceImpl {
	return &districtServiceImpl{
		districtRepository: districtRepository,
		villageRepository:  villageRepository,
	}
}

func (d *districtServiceImpl) GetAll(ctx context.Context, keyword string) (responses []model.District, err error) {
	var districts []entity.District
	var repoErr error

	if keyword == "" {
		districts, repoErr = d.districtRepository.FindAll(ctx)
	} else {
		districts, repoErr = d.districtRepository.FindByName(ctx, keyword)
	}

	if repoErr != nil {
		err = mapError(repoErr)
		return
	}

	responses = make([]model.District, len(districts))

	for i, district := range districts {
		responses[i] = d.mapToModel(district)
	}
	return
}

func (d *districtServiceImpl) GetByID(ctx context.Context, id string) (response model.District, err error) {
	district, repoErr := d.districtRepository.FindByID(ctx, id)
	if repoErr != nil {
		err = mapError(repoErr)
		return
	}

	response = d.mapToModel(district)
	return
}

func (d *districtServiceImpl) GetVillagesByDistrictID(ctx context.Context, id string) (responses []model.Village, err error) {
	villages, repoErr := d.villageRepository.FindByDistrictID(ctx, id)
	if repoErr != nil {
		err = mapError(repoErr)
		return
	}

	responses = d.mapToModels(villages)
	return
}

func (d *districtServiceImpl) GetVillagesByDistrictName(ctx context.Context, keyword string) (responses []model.Village, err error) {
	villages, repoErr := d.villageRepository.FindByDistrictName(ctx, keyword)
	if repoErr != nil {
		err = mapError(repoErr)
		return
	}

	responses = d.mapToModels(villages)
	return
}

func (d *districtServiceImpl) mapToModel(e entity.District) model.District {
	return model.District{
		ID:   e.ID,
		Name: e.Name,
		Regency: model.Regency{
			ID:   e.Regency.ID,
			Name: e.Regency.Name,
			Province: model.Province{
				ID:   e.Regency.Province.ID,
				Name: e.Regency.Province.Name,
			},
		},
	}
}

func (d *districtServiceImpl) mapToModels(entities []entity.Village) []model.Village {
	villages := make([]model.Village, len(entities))

	for i, e := range entities {
		district := d.mapToModel(e.District)
		village := model.Village{
			ID:       e.ID,
			Name:     e.Name,
			District: district,
		}

		villages[i] = village
	}

	return villages
}
