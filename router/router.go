package router

import (
	"github.com/alexaxms/TopSecret/controller"
	"github.com/labstack/echo"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.GET("/boards", func(context echo.Context) error { return c.GetBoards(context) })

	return e
}
