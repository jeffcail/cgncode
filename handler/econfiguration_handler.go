package handler

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/jeffcail/cgncode/models"
    "github.com/jeffcail/cgncode/service"
)

type EConfigurationHandler struct{}

// CreateEConfiguration creates a new EConfiguration
func (this *EConfigurationHandler) CreateEConfiguration(c echo.Context) error {
    var input models.EConfiguration
    if err := c.Bind(&input); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }

    if err := service.NewEConfigurationService().CreateEConfiguration(&input); err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }

    return c.JSON(http.StatusOK, input)
}

// GetEConfiguration retrieves an existing EConfiguration
func (this *EConfigurationHandler) GetEConfiguration(c echo.Context) error {
    id := c.Param("id")

    result, err := service.NewEConfigurationService().GetEConfiguration(id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }

    return c.JSON(http.StatusOK, result)
}

// UpdateEConfiguration updates an existing EConfiguration
func (this *EConfigurationHandler) UpdateEConfiguration(c echo.Context) error {
    id := c.Param("id")
    var input models.EConfiguration
    if err := c.Bind(&input); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }

    if err := service.NewEConfigurationService().UpdateEConfiguration(id, &input); err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }

    return c.JSON(http.StatusOK, input)
}

// DeleteEConfiguration deletes an existing EConfiguration
func (this *EConfigurationHandler) DeleteEConfiguration(c echo.Context) error {
    id := c.Param("id")

    if err := service.NewEConfigurationService().DeleteEConfiguration(id); err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }

    return c.JSON(http.StatusOK, id)
}
