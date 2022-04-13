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

func TestVillagesController(t *testing.T) {
	t.Run("TestNewVillagesController", func(t *testing.T) {
		mockService := &mocks.VillageService{}
		controller := NewVillagesController(mockService)
		assert.NotNil(t, controller)
	})

	t.Run("TestRoute", func(t *testing.T) {
		mockService := &mocks.VillageService{}
		controller := NewVillagesController(mockService)
		g := echo.New().Group("/api/v1")
		controller.Route(g)
		assert.NotNil(t, controller)
	})

	t.Run("TestGetAll", func(t *testing.T) {
		mockService := &mocks.VillageService{}

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
			mockService.On("GetAll", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, keyword string) []model.Village {
					return dummyVillages
				},
				func(ctx context.Context, keyword string) error {
					return nil
				},
			).Once()

			t.Run("it should return 200 status code with valid response, when there is no error", func(t *testing.T) {
				controller := NewVillagesController(mockService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/villages?keyword=Pager", nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				if assert.NoError(t, controller.getAll(c)) {
					assert.Equal(t, http.StatusOK, rec.Code)

					body := rec.Body.String()

					response := make(map[string]interface{})
					if assert.NoError(t, json.Unmarshal([]byte(body), &response)) {
						status := response["status"].(string)
						message := response["message"].(string)
						data := response["data"].(map[string]any)["villages"].([]any)

						assert.Equal(t, "success", status)
						assert.Equal(t, "successfully get villages", message)
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
			mockService.On("GetAll", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, keyword string) []model.Village {
					return []model.Village{}
				},
				func(ctx context.Context, keyword string) error {
					return service.ErrRepository
				},
			).Once()

			t.Run("it should return 500 status code with valid response, when error happened", func(t *testing.T) {
				controller := NewVillagesController(mockService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/villages?keyword=Bungkal", nil)
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
		mockService := &mocks.VillageService{}

		dummyVillage := model.Village{
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
		}

		t.Run("success scenario", func(t *testing.T) {
			mockService.On("GetByID", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), dummyVillage.ID).Return(
				func(ctx context.Context, id string) model.Village {
					return dummyVillage
				},
				func(ctx context.Context, id string) error {
					return nil
				},
			).Once()

			t.Run("it should return 200 status code with valid response, when there is no error", func(t *testing.T) {
				controller := NewVillagesController(mockService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/villages", nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/:id")
				c.SetParamNames("id")
				c.SetParamValues(dummyVillage.ID)

				if assert.NoError(t, controller.getByID(c)) {
					assert.Equal(t, http.StatusOK, rec.Code)

					body := rec.Body.String()

					response := make(map[string]interface{})
					if assert.NoError(t, json.Unmarshal([]byte(body), &response)) {
						status := response["status"].(string)
						message := response["message"].(string)
						data := response["data"].(map[string]any)

						assert.Equal(t, "success", status)
						assert.Equal(t, fmt.Sprintf("successfully get village with ID %s", dummyVillage.ID), message)
						assert.Equal(t, dummyVillage.ID, data["id"])
						assert.Equal(t, dummyVillage.Name, data["name"])
						assert.Equal(t, dummyVillage.District.Regency.ID, data["district"].(map[string]any)["regency"].(map[string]any)["id"])
						assert.Equal(t, dummyVillage.District.Regency.Name, data["district"].(map[string]any)["regency"].(map[string]any)["name"])
						assert.Equal(t, dummyVillage.District.Regency.ID, data["district"].(map[string]any)["regency"].(map[string]any)["id"])
						assert.Equal(t, dummyVillage.District.Regency.Name, data["district"].(map[string]any)["regency"].(map[string]any)["name"])
						assert.Equal(t, dummyVillage.District.Regency.Province.ID, data["district"].(map[string]any)["regency"].(map[string]any)["province"].(map[string]any)["id"])
						assert.Equal(t, dummyVillage.District.Regency.Province.Name, data["district"].(map[string]any)["regency"].(map[string]any)["province"].(map[string]any)["name"])
					}
				}
			})
		})

		t.Run("failed scenario", func(t *testing.T) {
			mockService.On("GetByID", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
				func(ctx context.Context, id string) model.Village {
					return model.Village{}
				},
				func(ctx context.Context, id string) error {
					if id != dummyVillage.ID {
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
					id:                 dummyVillage.ID + "1",
					expectedStatusCode: http.StatusNotFound,
					expectedMessage:    "Resource with given ID not found.",
				},
				{
					name:               "it should return 500 status code with valid response, when error happened",
					id:                 dummyVillage.ID,
					expectedStatusCode: http.StatusInternalServerError,
					expectedMessage:    "Something went wrong.",
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					controller := NewVillagesController(mockService)

					e := echo.New()
					req := httptest.NewRequest(http.MethodGet, "/api/v1/villages", nil)
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
