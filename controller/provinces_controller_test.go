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

func TestProvincesController(t *testing.T) {
	t.Run("TestNewProvincesController", func(t *testing.T) {
		mockService := &mocks.ProvinceService{}
		controller := NewProvincesController(mockService)
		assert.NotNil(t, controller)
	})

	t.Run("TestRoute", func(t *testing.T) {
		mockService := &mocks.ProvinceService{}
		controller := NewProvincesController(mockService)
		g := echo.New().Group("/api/v1")
		controller.Route(g)
		assert.NotNil(t, controller)
	})

	t.Run("TestGetAll", func(t *testing.T) {
		mockService := &mocks.ProvinceService{}

		dummyProvinces := []model.Province{
			{
				ID:   "35",
				Name: "Jawa Timur",
			},
		}

		t.Run("success scenario", func(t *testing.T) {
			mockService.On("GetAll", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, keyword string) []model.Province {
					return dummyProvinces
				},
				func(ctx context.Context, keyword string) error {
					return nil
				},
			).Once()

			t.Run("it should return 200 status code with valid response, when there is no error", func(t *testing.T) {
				controller := NewProvincesController(mockService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/provinces?keyword=Jawa", nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				if assert.NoError(t, controller.getAll(c)) {
					assert.Equal(t, http.StatusOK, rec.Code)

					body := rec.Body.String()

					response := make(map[string]interface{})
					if assert.NoError(t, json.Unmarshal([]byte(body), &response)) {
						status := response["status"].(string)
						message := response["message"].(string)
						data := response["data"].(map[string]any)["provinces"].([]any)

						assert.Equal(t, "success", status)
						assert.Equal(t, "successfully get provinces", message)
						assert.Equal(t, len(dummyProvinces), len(data))

						for i, province := range dummyProvinces {
							gotProvince := data[i].(map[string]any)
							gotProvinceID := gotProvince["id"].(string)
							gotProvinceName := gotProvince["name"].(string)
							assert.Equal(t, province.ID, gotProvinceID)
							assert.Equal(t, province.Name, gotProvinceName)
						}
					}
				}
			})
		})

		t.Run("failed scenario", func(t *testing.T) {
			mockService.On("GetAll", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, keyword string) []model.Province {
					return []model.Province{}
				},
				func(ctx context.Context, keyword string) error {
					return service.ErrRepository
				},
			).Once()

			t.Run("it should return 500 status code with valid response, when error happened", func(t *testing.T) {
				controller := NewProvincesController(mockService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/provinces?keyword=Jawa", nil)
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
		mockService := &mocks.ProvinceService{}

		dummyProvince := model.Province{
			ID:   "35",
			Name: "Jawa Timur",
		}

		t.Run("success scenario", func(t *testing.T) {
			mockService.On("GetByID", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), dummyProvince.ID).Return(
				func(ctx context.Context, id string) model.Province {
					return dummyProvince
				},
				func(ctx context.Context, id string) error {
					return nil
				},
			).Once()

			t.Run("it should return 200 status code with valid response, when there is no error", func(t *testing.T) {
				controller := NewProvincesController(mockService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/provinces", nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/:id")
				c.SetParamNames("id")
				c.SetParamValues(dummyProvince.ID)

				if assert.NoError(t, controller.getByID(c)) {
					assert.Equal(t, http.StatusOK, rec.Code)

					body := rec.Body.String()

					response := make(map[string]interface{})
					if assert.NoError(t, json.Unmarshal([]byte(body), &response)) {
						status := response["status"].(string)
						message := response["message"].(string)
						data := response["data"].(map[string]any)

						assert.Equal(t, "success", status)
						assert.Equal(t, fmt.Sprintf("successfully get province with ID %s", data["id"]), message)
						assert.Equal(t, dummyProvince.ID, data["id"])
						assert.Equal(t, dummyProvince.Name, data["name"])
					}
				}
			})
		})

		t.Run("failed scenario", func(t *testing.T) {
			mockService.On("GetByID", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, id string) model.Province {
					return model.Province{}
				},
				func(ctx context.Context, id string) error {
					if id != dummyProvince.ID {
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
					id:                 dummyProvince.ID + "1",
					expectedStatusCode: http.StatusNotFound,
					expectedMessage:    "Resource with given ID not found.",
				},
				{
					name:               "it should return 500 status code with valid response, when error happened",
					id:                 dummyProvince.ID,
					expectedStatusCode: http.StatusInternalServerError,
					expectedMessage:    "Something went wrong.",
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					controller := NewProvincesController(mockService)

					e := echo.New()
					req := httptest.NewRequest(http.MethodGet, "/api/v1/provinces", nil)
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
