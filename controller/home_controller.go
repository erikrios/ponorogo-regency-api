package controller

import (
	"net/http"

	"github.com/erikrios/ponorogo-regency-api/model"
	"github.com/labstack/echo/v4"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (h *HomeController) Route(e *echo.Echo) {
	e.GET("/", h.getHello)
}

// GetHello godoc
// @Summary      Check the API Connectivity
// @Description  Show the hello message
// @Tags         home
// @Accept       json
// @Produce      json
// @Success      200  {object}  response
// @Router       / [get]
func (h *HomeController) getHello(c echo.Context) error {
	data := "Hello, World"
	return c.JSON(http.StatusOK, model.NewResponse("success", "successfully get message", &data))
}

// response struct used for swaggo to generate the API documentation, as it doesn't support generic yet.
type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
