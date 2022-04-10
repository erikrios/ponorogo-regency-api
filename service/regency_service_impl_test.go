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

func TestRegencyServiceImpl(t *testing.T) {

	t.Run("TestNewRegencyServiceImpl", func(t *testing.T) {
		mockRepo := &mocks.RegencyRepository{}

		t.Run("it should return valid regency service instance, when invoke the function", func(t *testing.T) {
			var service RegencyService = NewRegencyServiceImpl(mockRepo)
			assert.NotNil(t, service)
		})
	})

	t.Run("TestGetAll", func(t *testing.T) {
		mockRepo := &mocks.RegencyRepository{}

		dummyRegencies := []entity.Regency{
			{
				ID:   "5301",
				Name: "Ponorogo",
				Province: entity.Province{
					ID:   "53",
					Name: "Jawa Timur",
				},
			},
		}

		t.Run("success scenario", func(t *testing.T) {
			mockRepo.On("FindAll", mock.AnythingOfType(fmt.Sprintf("%T", context.Background()))).Return(
				func(ctx context.Context) []entity.Regency {
					return dummyRegencies
				},
				func(ctx context.Context) error {
					return nil
				},
			).Once()
			mockRepo.On("FindByName", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, keyword string) []entity.Regency {
					return dummyRegencies
				},
				func(ctx context.Context, keyword string) error {
					return nil
				},
			).Once()

			testCases := []struct {
				name     string
				keyword  string
				expected []model.Regency
			}{
				{
					name:     "it should return valid regencies, when keyword is not empty",
					keyword:  dummyRegencies[0].Name,
					expected: mapToRegenciesModel(dummyRegencies),
				},
				{
					name:     "it should return valid regencies, when keyword is empty",
					keyword:  "",
					expected: mapToRegenciesModel(dummyRegencies),
				},
			}

			var service RegencyService = NewRegencyServiceImpl(mockRepo)

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
				func(ctx context.Context) []entity.Regency {
					return []entity.Regency{}
				},
				func(ctx context.Context) error {
					return repository.ErrDatabase
				},
			).Once()
			mockRepo.On("FindByName", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, keyword string) []entity.Regency {
					return []entity.Regency{}
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
					keyword:  dummyRegencies[0].Name,
					expected: ErrRepository,
				},
				{
					name:     "it should return ErrRepository instance, when keyword is empty and error happened",
					keyword:  "",
					expected: ErrRepository,
				},
			}

			var service RegencyService = NewRegencyServiceImpl(mockRepo)

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					_, err := service.GetAll(context.Background(), testCase.keyword)
					assert.ErrorIs(t, err, testCase.expected)
				})
			}
		})
	})

	t.Run("TestGetByID", func(t *testing.T) {
		mockRepo := &mocks.RegencyRepository{}

		dummyRegency := entity.Regency{
			ID:   "3302",
			Name: "Ponorogo",
			Province: entity.Province{
				ID:   "33",
				Name: "Jawa Timur",
			},
		}

		t.Run("success scenario", func(t *testing.T) {
			mockRepo.On("FindByID", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, id string) entity.Regency {
					return dummyRegency
				},
				func(ctx context.Context, id string) error {
					return nil
				},
			).Once()

			testCases := []struct {
				name     string
				ID       string
				expected model.Regency
			}{
				{
					name:     "it should return valid regency, when ID is valid",
					ID:       dummyRegency.ID,
					expected: mapToRegencyModel(dummyRegency),
				},
			}

			var service RegencyService = NewRegencyServiceImpl(mockRepo)

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
				func(ctx context.Context, id string) entity.Regency {
					return entity.Regency{}
				},
				func(ctx context.Context, id string) error {
					if id != dummyRegency.ID {
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
					id:       dummyRegency.ID,
					expected: ErrRepository,
				},
				{
					name:     "it should return ErrDataNotFound instance, when ID is not match and error happened",
					id:       "9090",
					expected: ErrDataNotFound,
				},
			}

			var service RegencyService = NewRegencyServiceImpl(mockRepo)

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					_, err := service.GetByID(context.Background(), testCase.id)
					assert.ErrorIs(t, err, testCase.expected)
				})
			}
		})
	})
}

func mapToRegencyModel(e entity.Regency) model.Regency {
	return model.Regency{
		ID:   e.ID,
		Name: e.Name,
		Province: model.Province{
			ID:   e.Province.ID,
			Name: e.Province.Name,
		},
	}
}

func mapToRegenciesModel(entities []entity.Regency) []model.Regency {
	regencies := make([]model.Regency, len(entities))

	for i, e := range entities {
		regencies[i] = model.Regency{
			ID:   e.ID,
			Name: e.Name,
			Province: model.Province{
				ID:   e.Province.ID,
				Name: e.Province.Name,
			},
		}
	}

	return regencies
}
