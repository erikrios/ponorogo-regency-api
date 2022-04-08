package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/erikrios/ponorogo-regency-api/model"
	"github.com/erikrios/ponorogo-regency-api/service"
	"github.com/labstack/echo/v4"
)

type villagesController struct {
	service service.VillageService
}

func NewVillagesController(service service.VillageService) *villagesController {
	return &villagesController{service: service}
}

func (v *villagesController) Route(g *echo.Group) {
	group := g.Group("/villages")
	group.GET("", v.getAll)
	group.GET("/:id", v.getByID)
}

// GetAll	     godoc
// @Summary      Get Villages
// @Description  Get villages
// @Tags         villages
// @Accept       json
// @Produce      json
// @Param        keyword  query     string  false  "village name search by keyword"
// @Success      200      {object}  villagesResponse
// @Failure      500      {object}  echo.HTTPError
// @Router       /api/v1/villages [get]
func (v *villagesController) getAll(c echo.Context) error {
	keyword := c.QueryParam("keyword")

	villages, err := v.service.GetAll(context.Background(), keyword)
	if err != nil {
		return newErrorResponse(err)
	}

	response := model.NewResponse("success", "successfully get villages", villages)
	return c.JSON(http.StatusOK, response)
}

// GetByID       godoc
// @Summary      Get Village by ID
// @Description  get villages by ID
// @Tags         villages
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Village ID"
// @Success      200  {object}  villageResponse
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /api/v1/villages/{id} [get]
func (v *villagesController) getByID(c echo.Context) error {
	id := c.Param("id")

	village, err := v.service.GetByID(context.Background(), id)
	if err != nil {
		return newErrorResponse(err)
	}

	response := model.NewResponse("success", fmt.Sprintf("successfully get village with ID %s", id), village)
	return c.JSON(http.StatusOK, response)
}

// villagesResponse struct is used for swaggo to generate the API documentation, as it doesn't support generic yet.
type villagesResponse struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    []model.Village `json:"data"`
}

// villageResponse struct is used for swaggo to generate the API documentation, as it doesn't support generic yet.
type villageResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    model.Village `json:"data"`
}
