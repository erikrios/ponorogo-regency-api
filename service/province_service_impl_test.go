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

func TestProvinceServiceImpl(t *testing.T) {

	t.Run("TestNewProvinceServiceImpl", func(t *testing.T) {
		mockRepo := &mocks.ProvinceRepository{}

		t.Run("it should return valid province service instance, when invoke the function", func(t *testing.T) {
			var service ProvinceService = NewProvinceServiceImpl(mockRepo)
			assert.NotNil(t, service)
		})
	})

	t.Run("TestGetAll", func(t *testing.T) {
		mockRepo := &mocks.ProvinceRepository{}

		dummyProvinces := []entity.Province{
			{
				ID:   "33",
				Name: "Jawa Timur",
			},
		}

		t.Run("success scenario", func(t *testing.T) {
			mockRepo.On("FindAll", mock.AnythingOfType(fmt.Sprintf("%T", context.Background()))).Return(
				func(ctx context.Context) []entity.Province {
					return dummyProvinces
				},
				func(ctx context.Context) error {
					return nil
				},
			).Once()
			mockRepo.On("FindByName", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, keyword string) []entity.Province {
					return dummyProvinces
				},
				func(ctx context.Context, keyword string) error {
					return nil
				},
			).Once()

			testCases := []struct {
				name     string
				keyword  string
				expected []model.Province
			}{
				{
					name:     "it should return valid provinces, when keyword is not empty",
					keyword:  dummyProvinces[0].Name,
					expected: mapToProvincesModel(dummyProvinces),
				},
				{
					name:     "it should return valid provinces, when keyword is empty",
					keyword:  "",
					expected: mapToProvincesModel(dummyProvinces),
				},
			}

			var service ProvinceService = NewProvinceServiceImpl(mockRepo)

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
				func(ctx context.Context) []entity.Province {
					return []entity.Province{}
				},
				func(ctx context.Context) error {
					return repository.ErrDatabase
				},
			).Once()
			mockRepo.On("FindByName", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, keyword string) []entity.Province {
					return []entity.Province{}
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
					keyword:  dummyProvinces[0].Name,
					expected: ErrRepository,
				},
				{
					name:     "it should return ErrRepository instance, when keyword is empty and error happened",
					keyword:  "",
					expected: ErrRepository,
				},
			}

			var service ProvinceService = NewProvinceServiceImpl(mockRepo)

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					_, err := service.GetAll(context.Background(), testCase.keyword)
					assert.ErrorIs(t, err, testCase.expected)
				})
			}
		})
	})

	t.Run("TestGetByID", func(t *testing.T) {
		mockRepo := &mocks.ProvinceRepository{}

		dummyProvince := entity.Province{
			ID:   "33",
			Name: "Jawa Timur",
		}

		t.Run("success scenario", func(t *testing.T) {
			mockRepo.On("FindByID", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, id string) entity.Province {
					return dummyProvince
				},
				func(ctx context.Context, id string) error {
					return nil
				},
			).Once()

			testCases := []struct {
				name     string
				ID       string
				expected model.Province
			}{
				{
					name:     "it should return valid province, when ID is valid",
					ID:       dummyProvince.ID,
					expected: mapToProvinceModel(dummyProvince),
				},
			}

			var service ProvinceService = NewProvinceServiceImpl(mockRepo)

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
				func(ctx context.Context, id string) entity.Province {
					return entity.Province{}
				},
				func(ctx context.Context, id string) error {
					if id != dummyProvince.ID {
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
					id:       dummyProvince.ID,
					expected: ErrRepository,
				},
				{
					name:     "it should return ErrDataNotFound instance, when ID is not match and error happened",
					id:       "90",
					expected: ErrDataNotFound,
				},
			}

			var service ProvinceService = NewProvinceServiceImpl(mockRepo)

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					_, err := service.GetByID(context.Background(), testCase.id)
					assert.ErrorIs(t, err, testCase.expected)
				})
			}
		})
	})
}

func mapToProvinceModel(e entity.Province) model.Province {
	return model.Province{
		ID:   e.ID,
		Name: e.Name,
	}
}

func mapToProvincesModel(entities []entity.Province) []model.Province {
	provinces := make([]model.Province, len(entities))

	for i, e := range entities {
		provinces[i] = mapToProvinceModel(e)
	}

	return provinces
}
