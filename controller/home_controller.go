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
	e.GET("/", h.GetHello)
}

func (h *HomeController) GetHello(c echo.Context) error {
	data := "Hello, World"
	return c.JSON(http.StatusOK, model.NewResponse("success", "successfully get message", &data))
}
