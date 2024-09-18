package handler

import (
	"github.com/jeffcail/cgncode/models"
	"github.com/jeffcail/cgncode/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type InfoHandler struct{}

// CreateInfo creates a new Info
func (this *InfoHandler) CreateInfo(c echo.Context) error {
	var input models.Info
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := service.NewInfoService().CreateInfo(&input); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, input)
}

// GetInfo retrieves an existing Info
func (this *InfoHandler) GetInfo(c echo.Context) error {
	id := c.Param("id")

	result, err := service.NewInfoService().GetInfo(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

// UpdateInfo updates an existing Info
func (this *InfoHandler) UpdateInfo(c echo.Context) error {
	id := c.Param("id")
	var input models.Info
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := service.NewInfoService().UpdateInfo(id, &input); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, input)
}

// DeleteInfo deletes an existing Info
func (this *InfoHandler) DeleteInfo(c echo.Context) error {
	id := c.Param("id")

	if err := service.NewInfoService().DeleteInfo(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, id)
}
