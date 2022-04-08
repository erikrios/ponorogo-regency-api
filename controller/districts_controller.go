package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/erikrios/ponorogo-regency-api/model"
	"github.com/erikrios/ponorogo-regency-api/service"
	"github.com/labstack/echo/v4"
)

type districtsController struct {
	service service.DistrictService
}

func NewDistrictsController(service service.DistrictService) *districtsController {
	return &districtsController{service: service}
}

func (d *districtsController) Route(g *echo.Group) {
	group := g.Group("/districts")
	group.GET("", d.getAll)
	group.GET("/:id", d.getByID)
	group.GET("/:id/villages", d.getVillagesByDistrictID)
	group.GET("/villages", d.getVillagesByDistrictName)
}

// GetAll	     godoc
// @Summary      Get Districts
// @Description  Get districts
// @Tags         districts
// @Accept       json
// @Produce      json
// @Param        keyword  query     string  false  "district name search by keyword"
// @Success      200      {object}  districtsResponse
// @Failure      500  {object}  echo.HTTPError
// @Router       /api/v1/districts [get]
func (d *districtsController) getAll(c echo.Context) error {
	keyword := c.QueryParam("keyword")

	districts, err := d.service.GetAll(context.Background(), keyword)
	if err != nil {
		return newErrorResponse(err)
	}

	response := model.NewResponse("success", "successfully get districts", districts)
	return c.JSON(http.StatusOK, response)
}

// GetByID       godoc
// @Summary      Get District by ID
// @Description  get districts by ID
// @Tags         districts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "District ID"
// @Success      200  {object}  districtResponse
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /api/v1/districts/{id} [get]
func (p *districtsController) getByID(c echo.Context) error {
	id := c.Param("id")

	district, err := p.service.GetByID(context.Background(), id)
	if err != nil {
		return newErrorResponse(err)
	}

	response := model.NewResponse("success", fmt.Sprintf("successfully get district with ID %s", id), district)
	return c.JSON(http.StatusOK, response)
}

// GetVillagesByDistrictID	     godoc
// @Summary      Get Villages by District ID
// @Description  Get villages by district ID
// @Tags         districts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "District ID"
// @Success      200  {object}  villagesResponse
// @Failure      500      {object}  echo.HTTPError
// @Router       /api/v1/districts/{id}/villages [get]
func (p *districtsController) getVillagesByDistrictID(c echo.Context) error {
	id := c.Param("id")

	villages, err := p.service.GetVillagesByDistrictID(context.Background(), id)
	if err != nil {
		return newErrorResponse(err)
	}

	response := model.NewResponse("success", fmt.Sprintf("successfully get villages with district ID %s", id), villages)
	return c.JSON(http.StatusOK, response)
}

// GetVillagesByDistrictName     godoc
// @Summary      Get Villages by District Name
// @Description  Get villages by district name
// @Tags         districts
// @Accept       json
// @Produce      json
// @Param        keyword  query     string  false  "district name search by keyword"
// @Success      200      {object}  villagesResponse
// @Failure      500      {object}  echo.HTTPError
// @Router       /api/v1/districts/villages [get]
func (p *districtsController) getVillagesByDistrictName(c echo.Context) error {
	keyword := c.QueryParam("keyword")

	villages, err := p.service.GetVillagesByDistrictName(context.Background(), keyword)
	if err != nil {
		return newErrorResponse(err)
	}

	response := model.NewResponse("success", fmt.Sprintf("successfully get villages with district keyword name %s", keyword), villages)
	return c.JSON(http.StatusOK, response)
}

// districtsResponse struct is used for swaggo to generate the API documentation, as it doesn't support generic yet.
type districtsResponse struct {
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Data    []model.District `json:"data"`
}

// districtResponse struct is used for swaggo to generate the API documentation, as it doesn't support generic yet.
type districtResponse struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Data    model.District `json:"data"`
}
