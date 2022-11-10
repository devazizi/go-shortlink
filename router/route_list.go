package router

import (
	"goshorter/adapter"
	deliveryHttp "goshorter/delivery/http"

	"github.com/labstack/echo/v4"
)

func RegisterRouteList(router *echo.Echo, db adapter.DB) {
	api := router.Group("/api")
	{
		api.POST("/create-short-link", deliveryHttp.RegisterNewRoute(db))

	}
	router.GET("/r/:id", deliveryHttp.RedirectToRealLink(db))
}
