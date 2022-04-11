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

func TestVillageServiceImpl(t *testing.T) {

	t.Run("TestNewVillageServiceImpl", func(t *testing.T) {
		mockRepo := &mocks.VillageRepository{}

		t.Run("it should return valid village service instance, when invoke the function", func(t *testing.T) {
			var service VillageService = NewVillageServiceImpl(mockRepo)
			assert.NotNil(t, service)
		})
	})

	t.Run("TestGetAll", func(t *testing.T) {
		mockRepo := &mocks.VillageRepository{}

		dummyVillages := []entity.Village{
			{
				ID:   "53000101",
				Name: "Pager",
				District: entity.District{
					ID:   "530101",
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
			mockRepo.On("FindAll", mock.AnythingOfType(fmt.Sprintf("%T", context.Background()))).Return(
				func(ctx context.Context) []entity.Village {
					return dummyVillages
				},
				func(ctx context.Context) error {
					return nil
				},
			).Once()
			mockRepo.On("FindByName", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
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
					name:     "it should return valid villages, when keyword is not empty",
					keyword:  dummyVillages[0].Name,
					expected: mapToVillagesModel(dummyVillages),
				},
				{
					name:     "it should return valid regencies, when keyword is empty",
					keyword:  "",
					expected: mapToVillagesModel(dummyVillages),
				},
			}

			var service VillageService = NewVillageServiceImpl(mockRepo)

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					got, err := service.GetAll(context.Background(), testCase.keyword)
					assert.NoError(t, err)
					assert.ElementsMatch(t, testCase.expected, got)
				})
			}
		})

		t.Run("failed scenario", func(t *testing.T) {
			mockRepo.On("FindAll", mock.AnythingOfType(fmt.Sprintf("%T", context.Background()))).Return(
				func(ctx context.Context) []entity.Village {
					return []entity.Village{}
				},
				func(ctx context.Context) error {
					return repository.ErrDatabase
				},
			).Once()
			mockRepo.On("FindByName", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
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
					name:     "it should return ErrRepository instance, when keyword is not empty and error happened",
					keyword:  dummyVillages[0].Name,
					expected: ErrRepository,
				},
				{
					name:     "it should return ErrRepository instance, when keyword is empty and error happened",
					keyword:  "",
					expected: ErrRepository,
				},
			}

			var service VillageService = NewVillageServiceImpl(mockRepo)

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					_, err := service.GetAll(context.Background(), testCase.keyword)
					assert.ErrorIs(t, err, testCase.expected)
				})
			}
		})
	})

	t.Run("TestGetByID", func(t *testing.T) {
		mockRepo := &mocks.VillageRepository{}

		dummyVillage := entity.Village{
			ID:   "53000101",
			Name: "Pager",
			District: entity.District{
				ID:   "530101",
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
			mockRepo.On("FindByID", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, id string) entity.Village {
					return dummyVillage
				},
				func(ctx context.Context, id string) error {
					return nil
				},
			).Once()

			testCases := []struct {
				name     string
				ID       string
				expected model.Village
			}{
				{
					name:     "it should return valid village, when ID is valid",
					ID:       dummyVillage.ID,
					expected: mapToVillageModel(dummyVillage),
				},
			}

			var service VillageService = NewVillageServiceImpl(mockRepo)

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					got, err := service.GetByID(context.Background(), testCase.ID)
					assert.NoError(t, err)
					assert.Equal(t, testCase.expected, got)
				})
			}

		})

		t.Run("failed scenario", func(t *testing.T) {
			mockRepo.On("FindByID", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, id string) entity.Village {
					return entity.Village{}
				},
				func(ctx context.Context, id string) error {
					if id != dummyVillage.ID {
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
					id:       dummyVillage.ID,
					expected: ErrRepository,
				},
				{
					name:     "it should return ErrDataNotFound instance, when ID is not match and error happened",
					id:       "9090",
					expected: ErrDataNotFound,
				},
			}

			var service VillageService = NewVillageServiceImpl(mockRepo)

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					_, err := service.GetByID(context.Background(), testCase.id)
					assert.ErrorIs(t, err, testCase.expected)
				})
			}
		})
	})
}

func mapToVillageModel(e entity.Village) model.Village {
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

func mapToVillagesModel(entities []entity.Village) []model.Village {
	villages := make([]model.Village, len(entities))

	for i, e := range entities {
		villages[i] = mapToVillageModel(e)
	}

	return villages
}
