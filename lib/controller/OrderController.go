package controller

import (
	"net/http"

	"github.com/YashwantSingh7062/service2/lib/model"
	"github.com/labstack/echo"
)

func (c Controller) GetOrder(ctx echo.Context) error {
	order := model.Order{
		Id:        "123",
		Type:      "Cash",
		CreatedAt: "today",
	}
	return ctx.JSON(http.StatusOK, order)
}
