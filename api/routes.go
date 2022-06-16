package api

import (
	"github.com/labstack/echo"
)

func (a *API) RegisterRoutes(e *echo.Echo) {
	e.GET("/calendario", a.getCalendario)
	e.GET("/calendarios", a.getCalendarios)
	e.POST("/calendario", a.postCalendario)
	e.PATCH("/calendario", a.patchCalendario)
	e.DELETE("/calendario", a.deleteCalendario)
}
