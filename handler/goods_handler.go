package handler

import (
	"github.com/jeffcail/cgncode/models"
	"github.com/jeffcail/cgncode/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GoodsHandler struct{}

// CreateGoods creates a new Goods
func (this *GoodsHandler) CreateGoods(c echo.Context) error {
	var input models.Goods
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := service.NewGoodsService().CreateGoods(&input); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, input)
}

// GetGoods retrieves an existing Goods
func (this *GoodsHandler) GetGoods(c echo.Context) error {
	id := c.Param("id")

	result, err := service.NewGoodsService().GetGoods(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

// UpdateGoods updates an existing Goods
func (this *GoodsHandler) UpdateGoods(c echo.Context) error {
	id := c.Param("id")
	var input models.Goods
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := service.NewGoodsService().UpdateGoods(id, &input); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, input)
}

// DeleteGoods deletes an existing Goods
func (this *GoodsHandler) DeleteGoods(c echo.Context) error {
	id := c.Param("id")

	if err := service.NewGoodsService().DeleteGoods(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, id)
}
