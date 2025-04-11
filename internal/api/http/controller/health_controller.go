package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HelathController struct {
}

func NewHealthController() *HelathController {
	return &HelathController{}
}

func (hc *HelathController) Check(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
