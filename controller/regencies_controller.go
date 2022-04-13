package controller

import (
	"fmt"
	"net/http"

	"github.com/erikrios/ponorogo-regency-api/model"
	"github.com/erikrios/ponorogo-regency-api/service"
	"github.com/labstack/echo/v4"
)

type regenciesController struct {
	service service.RegencyService
}

func NewRegenciesController(service service.RegencyService) *regenciesController {
	return &regenciesController{service: service}
}

func (r *regenciesController) Route(g *echo.Group) {
	group := g.Group("/regencies")
	group.GET("", r.getAll)
	group.GET("/:id", r.getByID)
}

// GetAll	     	 godoc
// @Summary      Get Regencies
// @Description  Get regencies
// @Tags         regencies
// @Accept       json
// @Produce      json
// @Param        keyword  query     string  false  "regency name search by keyword"
// @Success      200      {object}  regenciesResponse
// @Failure      500      {object}  echo.HTTPError
// @Router       /regencies [get]
func (r *regenciesController) getAll(c echo.Context) error {
	keyword := c.QueryParam("keyword")

	regencies, err := r.service.GetAll(c.Request().Context(), keyword)
	if err != nil {
		return newErrorResponse(err)
	}

	regenciesResponse := map[string]any{"regencies": regencies}

	response := model.NewResponse("success", "successfully get regencies", regenciesResponse)
	return c.JSON(http.StatusOK, response)
}

// GetByID       godoc
// @Summary      Get Regency by ID
// @Description  get regencies by ID
// @Tags         regencies
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Regency ID"
// @Success      200  {object}  regencyResponse
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /regencies/{id} [get]
func (r *regenciesController) getByID(c echo.Context) error {
	id := c.Param("id")

	regency, err := r.service.GetByID(c.Request().Context(), id)
	if err != nil {
		return newErrorResponse(err)
	}

	response := model.NewResponse("success", fmt.Sprintf("successfully get regency with ID %s", id), regency)
	return c.JSON(http.StatusOK, response)
}

// regenciesResponse struct is used for swaggo to generate the API documentation, as it doesn't support generic yet.
type regenciesResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    regenciesData `json:"data"`
}

type regenciesData struct {
	Regencies []model.Regency `json:"regencies"`
}

// regencyResponse struct is used for swaggo to generate the API documentation, as it doesn't support generic yet.
type regencyResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    model.Regency `json:"data"`
}
