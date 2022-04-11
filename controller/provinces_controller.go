package controller

import (
	"fmt"
	"net/http"

	"github.com/erikrios/ponorogo-regency-api/model"
	"github.com/erikrios/ponorogo-regency-api/service"
	"github.com/labstack/echo/v4"
)

type provincesController struct {
	service service.ProvinceService
}

func NewProvincesController(service service.ProvinceService) *provincesController {
	return &provincesController{service: service}
}

func (p *provincesController) Route(g *echo.Group) {
	group := g.Group("/provinces")
	group.GET("", p.getAll)
	group.GET("/:id", p.getByID)
}

// GetAll	       godoc
// @Summary      Get Provinces
// @Description  Get provinces
// @Tags         provinces
// @Accept       json
// @Produce      json
// @Param        keyword  query     string  false  "province name search by keyword"
// @Success      200      {object}  provincesResponse
// @Failure      500      {object}  echo.HTTPError
// @Router       /provinces [get]
func (p *provincesController) getAll(c echo.Context) error {
	keyword := c.QueryParam("keyword")

	provinces, err := p.service.GetAll(c.Request().Context(), keyword)
	if err != nil {
		return newErrorResponse(err)
	}

	response := model.NewResponse("success", "successfully get provinces", provinces)
	return c.JSON(http.StatusOK, response)
}

// GetByID       godoc
// @Summary      Get Province by ID
// @Description  get provinces by ID
// @Tags         provinces
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Province ID"
// @Success      200  {object}  provinceResponse
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /provinces/{id} [get]
func (p *provincesController) getByID(c echo.Context) error {
	id := c.Param("id")

	province, err := p.service.GetByID(c.Request().Context(), id)
	if err != nil {
		return newErrorResponse(err)
	}

	response := model.NewResponse("success", fmt.Sprintf("successfully get province with ID %s", id), province)
	return c.JSON(http.StatusOK, response)
}

// provincesResponse struct is used for swaggo to generate the API documentation, as it doesn't support generic yet.
type provincesResponse struct {
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Data    []model.Province `json:"data"`
}

// provinceResponse struct is used for swaggo to generate the API documentation, as it doesn't support generic yet.
type provinceResponse struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Data    model.Province `json:"data"`
}
