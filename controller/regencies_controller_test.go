package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/erikrios/ponorogo-regency-api/model"
	"github.com/erikrios/ponorogo-regency-api/service"
	"github.com/erikrios/ponorogo-regency-api/service/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegenciesController(t *testing.T) {
	t.Run("TestNewRegenciesController", func(t *testing.T) {
		mockService := &mocks.RegencyService{}
		controller := NewRegenciesController(mockService)
		assert.NotNil(t, controller)
	})

	t.Run("TestRoute", func(t *testing.T) {
		mockService := &mocks.RegencyService{}
		controller := NewRegenciesController(mockService)
		g := echo.New().Group("/api/v1")
		controller.Route(g)
		assert.NotNil(t, controller)
	})

	t.Run("TestGetAll", func(t *testing.T) {
		mockService := &mocks.RegencyService{}

		dummyRegencies := []model.Regency{
			{
				ID:   "3501",
				Name: "Ponorogo",
				Province: model.Province{
					ID:   "35",
					Name: "Jawa Timur",
				},
			},
		}

		t.Run("success scenario", func(t *testing.T) {
			mockService.On("GetAll", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, keyword string) []model.Regency {
					return dummyRegencies
				},
				func(ctx context.Context, keyword string) error {
					return nil
				},
			).Once()

			t.Run("it should return 200 status code with valid response, when there is no error", func(t *testing.T) {
				controller := NewRegenciesController(mockService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/regencies?keyword=Ponorogo", nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				if assert.NoError(t, controller.getAll(c)) {
					assert.Equal(t, http.StatusOK, rec.Code)

					body := rec.Body.String()

					response := make(map[string]interface{})
					if assert.NoError(t, json.Unmarshal([]byte(body), &response)) {
						status := response["status"].(string)
						message := response["message"].(string)
						data := response["data"].([]any)

						assert.Equal(t, "success", status)
						assert.Equal(t, "successfully get regencies", message)
						assert.Equal(t, len(dummyRegencies), len(data))

						for i, province := range dummyRegencies {
							gotRegency := data[i].(map[string]any)
							gotRegencyID := gotRegency["id"].(string)
							gotRegencyName := gotRegency["name"].(string)
							assert.Equal(t, province.ID, gotRegencyID)
							assert.Equal(t, province.Name, gotRegencyName)
						}
					}
				}
			})
		})

		t.Run("failed scenario", func(t *testing.T) {
			mockService.On("GetAll", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, keyword string) []model.Regency {
					return []model.Regency{}
				},
				func(ctx context.Context, keyword string) error {
					return service.ErrRepository
				},
			).Once()

			t.Run("it should return 500 status code with valid response, when error happened", func(t *testing.T) {
				controller := NewRegenciesController(mockService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/provinces?keyword=Ponorogo", nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				gotError := controller.getAll(c)
				if assert.Error(t, gotError) {
					if echoHTTPError, ok := gotError.(*echo.HTTPError); assert.Equal(t, true, ok) {
						assert.Equal(t, http.StatusInternalServerError, echoHTTPError.Code)
						assert.Equal(t, "Something went wrong.", echoHTTPError.Message)
					}
				}
			})
		})
	})

	t.Run("TestGetByID", func(t *testing.T) {
		mockService := &mocks.RegencyService{}

		dummyRegency := model.Regency{
			ID:   "3501",
			Name: "Ponorogo",
			Province: model.Province{
				ID:   "35",
				Name: "Jawa Timur",
			},
		}

		t.Run("success scenario", func(t *testing.T) {
			mockService.On("GetByID", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), dummyRegency.ID).Return(
				func(ctx context.Context, id string) model.Regency {
					return dummyRegency
				},
				func(ctx context.Context, id string) error {
					return nil
				},
			).Once()

			t.Run("it should return 200 status code with valid response, when there is no error", func(t *testing.T) {
				controller := NewRegenciesController(mockService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/regencies", nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/:id")
				c.SetParamNames("id")
				c.SetParamValues(dummyRegency.ID)

				if assert.NoError(t, controller.getByID(c)) {
					assert.Equal(t, http.StatusOK, rec.Code)

					body := rec.Body.String()

					response := make(map[string]interface{})
					if assert.NoError(t, json.Unmarshal([]byte(body), &response)) {
						status := response["status"].(string)
						message := response["message"].(string)
						data := response["data"].(map[string]any)

						assert.Equal(t, "success", status)
						assert.Equal(t, fmt.Sprintf("successfully get regency with ID %s", data["id"]), message)
						assert.Equal(t, dummyRegency.ID, data["id"])
						assert.Equal(t, dummyRegency.Name, data["name"])
					}
				}
			})
		})

		t.Run("failed scenario", func(t *testing.T) {
			mockService.On("GetByID", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, id string) model.Regency {
					return model.Regency{}
				},
				func(ctx context.Context, id string) error {
					if id != dummyRegency.ID {
						return service.ErrDataNotFound
					}
					return service.ErrRepository
				},
			).Twice()

			testCases := []struct {
				name               string
				id                 string
				expectedStatusCode int
				expectedMessage    string
			}{
				{
					name:               "it should return 404 status code with valid response, when given ID not found",
					id:                 dummyRegency.ID + "1",
					expectedStatusCode: http.StatusNotFound,
					expectedMessage:    "Resource with given ID not found.",
				},
				{
					name:               "it should return 500 status code with valid response, when error happened",
					id:                 dummyRegency.ID,
					expectedStatusCode: http.StatusInternalServerError,
					expectedMessage:    "Something went wrong.",
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					controller := NewRegenciesController(mockService)

					e := echo.New()
					req := httptest.NewRequest(http.MethodGet, "/api/v1/regencies", nil)
					rec := httptest.NewRecorder()
					c := e.NewContext(req, rec)
					c.SetPath("/:id")
					c.SetParamNames("id")
					c.SetParamValues(testCase.id)

					gotError := controller.getByID(c)
					if assert.Error(t, gotError) {
						if echoHTTPError, ok := gotError.(*echo.HTTPError); assert.Equal(t, true, ok) {
							assert.Equal(t, testCase.expectedStatusCode, echoHTTPError.Code)
							assert.Equal(t, testCase.expectedMessage, echoHTTPError.Message)
						}
					}
				})
			}
		})
	})
}
