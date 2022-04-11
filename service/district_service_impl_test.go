package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/erikrios/ponorogo-regency-api/entity"
	"github.com/erikrios/ponorogo-regency-api/model"
	"github.com/erikrios/ponorogo-regency-api/repository"
	"github.com/erikrios/ponorogo-regency-api/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDistrictServiceImpl(t *testing.T) {

	t.Run("TestNewDistrictServiceImpl", func(t *testing.T) {
		mockDistrictRepo := &mocks.DistrictRepository{}
		mockVillageRepo := &mocks.VillageRepository{}

		t.Run("it should return valid district service instance, when invoke the function", func(t *testing.T) {
			var service DistrictService = NewDistrictServiceImpl(mockDistrictRepo, mockVillageRepo)
			assert.NotNil(t, service)
		})
	})

	t.Run("TestGetAll", func(t *testing.T) {
		mockDistrictRepo := &mocks.DistrictRepository{}
		mockVillageRepo := &mocks.VillageRepository{}

		dummyDistricts := []entity.District{
			{
				ID:   "35010000",
				Name: "Bungkal",
				Regency: entity.Regency{
					ID:   "5301",
					Name: "Ponorogo",
					Province: entity.Province{
						ID:   "53",
						Name: "Jawa Timur",
					},
				},
			},
		}

		t.Run("success scenario", func(t *testing.T) {
			mockDistrictRepo.On("FindAll", mock.AnythingOfType(fmt.Sprintf("%T", context.Background()))).Return(
				func(ctx context.Context) []entity.District {
					return dummyDistricts
				},
				func(ctx context.Context) error {
					return nil
				},
			).Once()
			mockDistrictRepo.On("FindByName", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, keyword string) []entity.District {
					return dummyDistricts
				},
				func(ctx context.Context, keyword string) error {
					return nil
				},
			).Once()

			testCases := []struct {
				name     string
				keyword  string
				expected []model.District
			}{
				{
					name:     "it should return valid districts, when keyword is not empty",
					keyword:  dummyDistricts[0].Name,
					expected: mapToDistrictsModel(dummyDistricts),
				},
				{
					name:     "it should return valid districts, when keyword is empty",
					keyword:  "",
					expected: mapToDistrictsModel(dummyDistricts),
				},
			}

			var service DistrictService = NewDistrictServiceImpl(mockDistrictRepo, mockVillageRepo)

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					got, err := service.GetAll(context.Background(), testCase.keyword)
					assert.NoError(t, err)
					assert.ElementsMatch(t, testCase.expected, got)
				})
			}
		})

		t.Run("failed scenario", func(t *testing.T) {
			mockDistrictRepo.On("FindAll", mock.AnythingOfType(fmt.Sprintf("%T", context.Background()))).Return(
				func(ctx context.Context) []entity.District {
					return []entity.District{}
				},
				func(ctx context.Context) error {
					return repository.ErrDatabase
				},
			).Once()
			mockDistrictRepo.On("FindByName", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, keyword string) []entity.District {
					return []entity.District{}
				},
				func(ctx context.Context, keyword string) error {
					return repository.ErrDatabase
				},
			).Once()

			testCases := []struct {
				name     string
				keyword  string
				expected error
			}{
				{
					name:     "it should return ErrRepository instance, when keyword is not empty and error happened",
					keyword:  dummyDistricts[0].Name,
					expected: ErrRepository,
				},
				{
					name:     "it should return ErrRepository instance, when keyword is empty and error happened",
					keyword:  "",
					expected: ErrRepository,
				},
			}

			var service DistrictService = NewDistrictServiceImpl(mockDistrictRepo, mockVillageRepo)

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					_, err := service.GetAll(context.Background(), testCase.keyword)
					assert.ErrorIs(t, err, testCase.expected)
				})
			}
		})
	})

	t.Run("TestGetByID", func(t *testing.T) {
		mockDistrictRepo := &mocks.DistrictRepository{}
		mockVillageRepo := &mocks.VillageRepository{}

		dummyDistrict := entity.District{
			ID:   "35200101",
			Name: "Bungkal",
			Regency: entity.Regency{
				ID:   "3302",
				Name: "Ponorogo",
				Province: entity.Province{
					ID:   "33",
					Name: "Jawa Timur",
				},
			},
		}

		t.Run("success scenario", func(t *testing.T) {
			mockDistrictRepo.On("FindByID", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, id string) entity.District {
					return dummyDistrict
				},
				func(ctx context.Context, id string) error {
					return nil
				},
			).Once()

			testCases := []struct {
				name     string
				ID       string
				expected model.District
			}{
				{
					name:     "it should return valid district, when ID is valid",
					ID:       dummyDistrict.ID,
					expected: mapToDistrictModel(dummyDistrict),
				},
			}

			var service DistrictService = NewDistrictServiceImpl(mockDistrictRepo, mockVillageRepo)

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					got, err := service.GetByID(context.Background(), testCase.ID)
					assert.NoError(t, err)
					assert.Equal(t, testCase.expected, got)
				})
			}
		})

		t.Run("failed scenario", func(t *testing.T) {
			mockDistrictRepo.On("FindByID", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, id string) entity.District {
					return entity.District{}
				},
				func(ctx context.Context, id string) error {
					if id != dummyDistrict.ID {
						return repository.ErrQueryNotFound
					} else {
						return repository.ErrDatabase
					}
				},
			).Twice()

			testCases := []struct {
				name     string
				id       string
				expected error
			}{
				{
					name:     "it should return ErrRepository instance, when ID is match and error happened",
					id:       dummyDistrict.ID,
					expected: ErrRepository,
				},
				{
					name:     "it should return ErrDataNotFound instance, when ID is not match and error happened",
					id:       "9090",
					expected: ErrDataNotFound,
				},
			}

			var service DistrictService = NewDistrictServiceImpl(mockDistrictRepo, mockVillageRepo)

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					_, err := service.GetByID(context.Background(), testCase.id)
					assert.ErrorIs(t, err, testCase.expected)
				})
			}
		})
	})

	t.Run("TestGetVillagesByDistrictID", func(t *testing.T) {
		mockDistrictRepo := &mocks.DistrictRepository{}
		mockVillageRepo := &mocks.VillageRepository{}

		dummyVillages := []entity.Village{
			{
				ID:   "35200101",
				Name: "Pager",
				District: entity.District{
					ID:   "350100",
					Name: "Bungkal",
					Regency: entity.Regency{
						ID:   "5301",
						Name: "Ponorogo",
						Province: entity.Province{
							ID:   "53",
							Name: "Jawa Timur",
						},
					},
				},
			},
		}

		t.Run("success scenario", func(t *testing.T) {
			mockVillageRepo.On("FindByDistrictID", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), dummyVillages[0].ID).Return(
				func(ctx context.Context, id string) []entity.Village {
					return dummyVillages
				},
				func(ctx context.Context, id string) error {
					return nil
				},
			).Once()

			testCases := []struct {
				name       string
				districtID string
				expected   []model.Village
			}{
				{
					name:       "it should return valid villages, when districtID is valid",
					districtID: dummyVillages[0].ID,
					expected:   mapToVillagesModel(dummyVillages),
				},
			}

			var service DistrictService = NewDistrictServiceImpl(mockDistrictRepo, mockVillageRepo)

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					got, err := service.GetVillagesByDistrictID(context.Background(), testCase.districtID)
					assert.NoError(t, err)
					assert.ElementsMatch(t, testCase.expected, got)
				})
			}
		})

		t.Run("failed scenario", func(t *testing.T) {
			mockVillageRepo.On("FindByDistrictID", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), dummyVillages[0].ID).Return(
				func(ctx context.Context, id string) []entity.Village {
					return []entity.Village{}
				},
				func(ctx context.Context, id string) error {
					return repository.ErrDatabase
				},
			).Once()

			testCases := []struct {
				name       string
				districtID string
				expected   error
			}{
				{
					name:       "it should return ErrRepository instance, when error happened",
					districtID: dummyVillages[0].ID,
					expected:   ErrRepository,
				},
			}

			var service DistrictService = NewDistrictServiceImpl(mockDistrictRepo, mockVillageRepo)

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					_, err := service.GetVillagesByDistrictID(context.Background(), testCase.districtID)
					assert.ErrorIs(t, err, testCase.expected)
				})
			}
		})
	})

	t.Run("TestGetVillagesByDistrictName", func(t *testing.T) {
		mockDistrictRepo := &mocks.DistrictRepository{}
		mockVillageRepo := &mocks.VillageRepository{}

		dummyVillages := []entity.Village{
			{
				ID:   "35200101",
				Name: "Pager",
				District: entity.District{
					ID:   "350100",
					Name: "Bungkal",
					Regency: entity.Regency{
						ID:   "5301",
						Name: "Ponorogo",
						Province: entity.Province{
							ID:   "53",
							Name: "Jawa Timur",
						},
					},
				},
			},
		}

		t.Run("success scenario", func(t *testing.T) {
			mockVillageRepo.On("FindByDistrictName", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), dummyVillages[0].Name).Return(
				func(ctx context.Context, keyword string) []entity.Village {
					return dummyVillages
				},
				func(ctx context.Context, keyword string) error {
					return nil
				},
			).Once()

			testCases := []struct {
				name     string
				keyword  string
				expected []model.Village
			}{
				{
					name:     "it should return valid villages, when keyword is valid",
					keyword:  dummyVillages[0].Name,
					expected: mapToVillagesModel(dummyVillages),
				},
			}

			var service DistrictService = NewDistrictServiceImpl(mockDistrictRepo, mockVillageRepo)

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					got, err := service.GetVillagesByDistrictName(context.Background(), testCase.keyword)
					assert.NoError(t, err)
					assert.ElementsMatch(t, testCase.expected, got)
				})
			}
		})

		t.Run("failed scenario", func(t *testing.T) {
			mockVillageRepo.On("FindByDistrictName", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), dummyVillages[0].Name).Return(
				func(ctx context.Context, keyword string) []entity.Village {
					return []entity.Village{}
				},
				func(ctx context.Context, keyword string) error {
					return repository.ErrDatabase
				},
			).Once()

			testCases := []struct {
				name     string
				keyword  string
				expected error
			}{
				{
					name:     "it should return ErrRepository instance, when error happened",
					keyword:  dummyVillages[0].Name,
					expected: ErrRepository,
				},
			}

			var service DistrictService = NewDistrictServiceImpl(mockDistrictRepo, mockVillageRepo)

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					_, err := service.GetVillagesByDistrictName(context.Background(), testCase.keyword)
					assert.ErrorIs(t, err, testCase.expected)
				})
			}
		})
	})
}

func mapToDistrictModel(e entity.District) model.District {
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

func mapToDistrictsModel(entities []entity.District) []model.District {
	districts := make([]model.District, len(entities))

	for i, e := range entities {
		districts[i] = mapToDistrictModel(e)
	}
	return districts
}
