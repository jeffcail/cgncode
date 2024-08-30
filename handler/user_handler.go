package handler

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/jeffcail/cgncode/models"
    "github.com/jeffcail/cgncode/service"
)

type UserHandler struct{}

// CreateUser creates a new User
func (this *UserHandler) CreateUser(c echo.Context) error {
    var input models.User
    if err := c.Bind(&input); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }

    if err := service.NewUserService().CreateUser(&input); err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }

    return c.JSON(http.StatusOK, input)
}

// GetUser retrieves an existing User
func (this *UserHandler) GetUser(c echo.Context) error {
    id := c.Param("id")

    result, err := service.NewUserService().GetUser(id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }

    return c.JSON(http.StatusOK, result)
}

// UpdateUser updates an existing User
func (this *UserHandler) UpdateUser(c echo.Context) error {
    id := c.Param("id")
    var input models.User
    if err := c.Bind(&input); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }

    if err := service.NewUserService().UpdateUser(id, &input); err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }

    return c.JSON(http.StatusOK, input)
}

// DeleteUser deletes an existing User
func (this *UserHandler) DeleteUser(c echo.Context) error {
    id := c.Param("id")

    if err := service.NewUserService().DeleteUser(id); err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }

    return c.JSON(http.StatusOK, id)
}
