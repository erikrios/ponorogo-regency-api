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

func TestDistrictsController(t *testing.T) {
	t.Run("TestNewDistrictsController", func(t *testing.T) {
		mockService := &mocks.DistrictService{}
		controller := NewDistrictsController(mockService)
		assert.NotNil(t, controller)
	})

	t.Run("TestRoute", func(t *testing.T) {
		mockService := &mocks.DistrictService{}
		controller := NewDistrictsController(mockService)
		g := echo.New().Group("/api/v1")
		controller.Route(g)
		assert.NotNil(t, controller)
	})

	t.Run("TestGetAll", func(t *testing.T) {
		mockService := &mocks.DistrictService{}

		dummyDistricts := []model.District{
			{
				ID:   "350101",
				Name: "Bungkal",
				Regency: model.Regency{
					ID:   "3501",
					Name: "Ponorogo",
					Province: model.Province{
						ID:   "35",
						Name: "Jawa Timur",
					},
				},
			},
		}

		t.Run("success scenario", func(t *testing.T) {
			mockService.On("GetAll", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, keyword string) []model.District {
					return dummyDistricts
				},
				func(ctx context.Context, keyword string) error {
					return nil
				},
			).Once()

			t.Run("it should return 200 status code with valid response, when there is no error", func(t *testing.T) {
				controller := NewDistrictsController(mockService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/districts?keyword=Bungkal", nil)
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
						assert.Equal(t, "successfully get districts", message)
						assert.Equal(t, len(dummyDistricts), len(data))

						for i, district := range dummyDistricts {
							gotDistrict := data[i].(map[string]any)
							gotDistrictID := gotDistrict["id"].(string)
							gotDistrictName := gotDistrict["name"].(string)
							gotRegencyID := gotDistrict["regency"].(map[string]any)["id"]
							gotRegencyName := gotDistrict["regency"].(map[string]any)["name"]
							gotProvinceID := gotDistrict["regency"].(map[string]any)["province"].(map[string]any)["id"]
							gotProvinceName := gotDistrict["regency"].(map[string]any)["province"].(map[string]any)["name"]
							assert.Equal(t, district.ID, gotDistrictID)
							assert.Equal(t, district.Name, gotDistrictName)
							assert.Equal(t, district.Regency.ID, gotRegencyID)
							assert.Equal(t, district.Regency.Name, gotRegencyName)
							assert.Equal(t, district.Regency.Province.ID, gotProvinceID)
							assert.Equal(t, district.Regency.Province.Name, gotProvinceName)
						}
					}
				}
			})
		})

		t.Run("failed scenario", func(t *testing.T) {
			mockService.On("GetAll", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, keyword string) []model.District {
					return []model.District{}
				},
				func(ctx context.Context, keyword string) error {
					return service.ErrRepository
				},
			).Once()

			t.Run("it should return 500 status code with valid response, when error happened", func(t *testing.T) {
				controller := NewDistrictsController(mockService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/districts?keyword=Bungkal", nil)
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
		mockService := &mocks.DistrictService{}

		dummyDistrict := model.District{
			ID:   "530101",
			Name: "Bungkal",
			Regency: model.Regency{
				ID:   "3501",
				Name: "Ponorogo",
				Province: model.Province{
					ID:   "35",
					Name: "Jawa Timur",
				},
			},
		}

		t.Run("success scenario", func(t *testing.T) {
			mockService.On("GetByID", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), dummyDistrict.ID).Return(
				func(ctx context.Context, id string) model.District {
					return dummyDistrict
				},
				func(ctx context.Context, id string) error {
					return nil
				},
			).Once()

			t.Run("it should return 200 status code with valid response, when there is no error", func(t *testing.T) {
				controller := NewDistrictsController(mockService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/districts", nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/:id")
				c.SetParamNames("id")
				c.SetParamValues(dummyDistrict.ID)

				if assert.NoError(t, controller.getByID(c)) {
					assert.Equal(t, http.StatusOK, rec.Code)

					body := rec.Body.String()

					response := make(map[string]interface{})
					if assert.NoError(t, json.Unmarshal([]byte(body), &response)) {
						status := response["status"].(string)
						message := response["message"].(string)
						data := response["data"].(map[string]any)

						assert.Equal(t, "success", status)
						assert.Equal(t, fmt.Sprintf("successfully get district with ID %s", dummyDistrict.ID), message)
						assert.Equal(t, dummyDistrict.ID, data["id"])
						assert.Equal(t, dummyDistrict.Name, data["name"])
						assert.Equal(t, dummyDistrict.Regency.ID, data["regency"].(map[string]any)["id"])
						assert.Equal(t, dummyDistrict.Regency.Name, data["regency"].(map[string]any)["name"])
						assert.Equal(t, dummyDistrict.Regency.Province.ID, data["regency"].(map[string]any)["province"].(map[string]any)["id"])
						assert.Equal(t, dummyDistrict.Regency.Province.Name, data["regency"].(map[string]any)["province"].(map[string]any)["name"])
					}
				}
			})
		})

		t.Run("failed scenario", func(t *testing.T) {
			mockService.On("GetByID", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, id string) model.District {
					return model.District{}
				},
				func(ctx context.Context, id string) error {
					if id != dummyDistrict.ID {
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
					id:                 dummyDistrict.ID + "1",
					expectedStatusCode: http.StatusNotFound,
					expectedMessage:    "Resource with given ID not found.",
				},
				{
					name:               "it should return 500 status code with valid response, when error happened",
					id:                 dummyDistrict.ID,
					expectedStatusCode: http.StatusInternalServerError,
					expectedMessage:    "Something went wrong.",
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					controller := NewDistrictsController(mockService)

					e := echo.New()
					req := httptest.NewRequest(http.MethodGet, "/api/v1/districts", nil)
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

	t.Run("TestGetVillagesByDistrictID", func(t *testing.T) {
		mockService := &mocks.DistrictService{}

		dummyVillages := []model.Village{
			{
				ID:   "35010101",
				Name: "Pager",
				District: model.District{
					ID:   "350101",
					Name: "Bungkal",
					Regency: model.Regency{
						ID:   "3501",
						Name: "Ponorogo",
						Province: model.Province{
							ID:   "35",
							Name: "Jawa Timur",
						},
					},
				},
			},
		}

		t.Run("success scenario", func(t *testing.T) {
			mockService.On("GetVillagesByDistrictID", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, districtID string) []model.Village {
					return dummyVillages
				},
				func(ctx context.Context, districtID string) error {
					return nil
				},
			).Once()

			t.Run("it should return 200 status code with valid response, when there is no error", func(t *testing.T) {
				controller := NewDistrictsController(mockService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/districts", nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/:id/villages")
				c.SetParamNames("id")
				c.SetParamValues(dummyVillages[0].District.ID)

				if assert.NoError(t, controller.getVillagesByDistrictID(c)) {
					assert.Equal(t, http.StatusOK, rec.Code)

					body := rec.Body.String()

					response := make(map[string]interface{})
					if assert.NoError(t, json.Unmarshal([]byte(body), &response)) {
						status := response["status"].(string)
						message := response["message"].(string)
						data := response["data"].([]any)

						assert.Equal(t, "success", status)
						assert.Equal(t, fmt.Sprintf("successfully get villages with district ID %s", dummyVillages[0].District.ID), message)
						assert.Equal(t, len(dummyVillages), len(data))

						for i, village := range dummyVillages {
							gotVillage := data[i].(map[string]any)
							gotVillageID := gotVillage["id"].(string)
							gotVillageName := gotVillage["name"].(string)
							gotDistrictID := gotVillage["district"].(map[string]any)["id"].(string)
							gotDistrictName := gotVillage["district"].(map[string]any)["name"].(string)
							gotRegencyID := gotVillage["district"].(map[string]any)["regency"].(map[string]any)["id"]
							gotRegencyName := gotVillage["district"].(map[string]any)["regency"].(map[string]any)["name"]
							gotProvinceID := gotVillage["district"].(map[string]any)["regency"].(map[string]any)["province"].(map[string]any)["id"]
							gotProvinceName := gotVillage["district"].(map[string]any)["regency"].(map[string]any)["province"].(map[string]any)["name"]
							assert.Equal(t, village.ID, gotVillageID)
							assert.Equal(t, village.Name, gotVillageName)
							assert.Equal(t, village.District.ID, gotDistrictID)
							assert.Equal(t, village.District.Name, gotDistrictName)
							assert.Equal(t, village.District.Regency.ID, gotRegencyID)
							assert.Equal(t, village.District.Regency.Name, gotRegencyName)
							assert.Equal(t, village.District.Regency.Province.ID, gotProvinceID)
							assert.Equal(t, village.District.Regency.Province.Name, gotProvinceName)
						}
					}
				}
			})
		})

		t.Run("failed scenario", func(t *testing.T) {
			mockService.On("GetVillagesByDistrictID", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, districtID string) []model.Village {
					return []model.Village{}
				},
				func(ctx context.Context, districtID string) error {
					return service.ErrRepository
				},
			).Once()

			t.Run("it should return 500 status code with valid response, when error happened", func(t *testing.T) {
				controller := NewDistrictsController(mockService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/districts", nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/:id/villages")
				c.SetParamNames("id")
				c.SetParamValues(dummyVillages[0].District.ID)

				gotError := controller.getVillagesByDistrictID(c)
				if assert.Error(t, gotError) {
					if echoHTTPError, ok := gotError.(*echo.HTTPError); assert.Equal(t, true, ok) {
						assert.Equal(t, http.StatusInternalServerError, echoHTTPError.Code)
						assert.Equal(t, "Something went wrong.", echoHTTPError.Message)
					}
				}
			})
		})
	})

	t.Run("TestGetVillagesByDistrictName", func(t *testing.T) {
		mockService := &mocks.DistrictService{}

		dummyVillages := []model.Village{
			{
				ID:   "35010101",
				Name: "Pager",
				District: model.District{
					ID:   "350101",
					Name: "Bungkal",
					Regency: model.Regency{
						ID:   "3501",
						Name: "Ponorogo",
						Province: model.Province{
							ID:   "35",
							Name: "Jawa Timur",
						},
					},
				},
			},
		}

		t.Run("success scenario", func(t *testing.T) {
			mockService.On("GetVillagesByDistrictName", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, keyword string) []model.Village {
					return dummyVillages
				},
				func(ctx context.Context, keyword string) error {
					return nil
				},
			).Once()

			t.Run("it should return 200 status code with valid response, when there is no error", func(t *testing.T) {
				controller := NewDistrictsController(mockService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/districts/villages?keyword=Bungkal", nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				if assert.NoError(t, controller.getVillagesByDistrictName(c)) {
					assert.Equal(t, http.StatusOK, rec.Code)

					body := rec.Body.String()

					response := make(map[string]interface{})
					if assert.NoError(t, json.Unmarshal([]byte(body), &response)) {
						status := response["status"].(string)
						message := response["message"].(string)
						data := response["data"].([]any)

						assert.Equal(t, "success", status)
						assert.Equal(t, fmt.Sprintf("successfully get villages with district keyword name %s", dummyVillages[0].District.Name), message)
						assert.Equal(t, len(dummyVillages), len(data))

						for i, village := range dummyVillages {
							gotVillage := data[i].(map[string]any)
							gotVillageID := gotVillage["id"].(string)
							gotVillageName := gotVillage["name"].(string)
							gotDistrictID := gotVillage["district"].(map[string]any)["id"].(string)
							gotDistrictName := gotVillage["district"].(map[string]any)["name"].(string)
							gotRegencyID := gotVillage["district"].(map[string]any)["regency"].(map[string]any)["id"]
							gotRegencyName := gotVillage["district"].(map[string]any)["regency"].(map[string]any)["name"]
							gotProvinceID := gotVillage["district"].(map[string]any)["regency"].(map[string]any)["province"].(map[string]any)["id"]
							gotProvinceName := gotVillage["district"].(map[string]any)["regency"].(map[string]any)["province"].(map[string]any)["name"]
							assert.Equal(t, village.ID, gotVillageID)
							assert.Equal(t, village.Name, gotVillageName)
							assert.Equal(t, village.District.ID, gotDistrictID)
							assert.Equal(t, village.District.Name, gotDistrictName)
							assert.Equal(t, village.District.Regency.ID, gotRegencyID)
							assert.Equal(t, village.District.Regency.Name, gotRegencyName)
							assert.Equal(t, village.District.Regency.Province.ID, gotProvinceID)
							assert.Equal(t, village.District.Regency.Province.Name, gotProvinceName)
						}
					}
				}
			})
		})

		t.Run("failed scenario", func(t *testing.T) {
			mockService.On("GetVillagesByDistrictName", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, keyword string) []model.Village {
					return []model.Village{}
				},
				func(ctx context.Context, keyword string) error {
					return service.ErrRepository
				},
			).Once()

			t.Run("it should return 500 status code with valid response, when error happened", func(t *testing.T) {
				controller := NewDistrictsController(mockService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/districts/villages?keyword=Bungkal", nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				gotError := controller.getVillagesByDistrictName(c)
				if assert.Error(t, gotError) {
					if echoHTTPError, ok := gotError.(*echo.HTTPError); assert.Equal(t, true, ok) {
						assert.Equal(t, http.StatusInternalServerError, echoHTTPError.Code)
						assert.Equal(t, "Something went wrong.", echoHTTPError.Message)
					}
				}
			})
		})
	})
}
