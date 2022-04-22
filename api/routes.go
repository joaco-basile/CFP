package api

import (
	"github.com/labstack/echo"
)

func (a *API) RegisterRoutes(e *echo.Echo) {
	e.POST("/calendario", a.postCalendario)
	e.GET("/calendario", a.getCalendario)
	e.PATCH("/calendario", a.patchCalendario)
	e.DELETE("/calendario", a.deleteCalendario)
}
