package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"use/internal/model"
	"use/internal/service"

	"github.com/labstack/echo/v4"
)

type UserAPI struct {
	userService *service.UserService
}

func NewUserAPI(userService *service.UserService) *UserAPI {
	return &UserAPI{userService: userService}
}

func (api *UserAPI) ListUsers(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	perPage, _ := strconv.Atoi(c.QueryParam("per_page"))
	if page <= 0 {
		page = 1
	}

	if perPage <= 0 {
		perPage = 10
	}

	users, err := api.userService.ListUsers(uint64(page), uint64(perPage))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

func (api *UserAPI) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	user, err := api.userService.GetUser(int64(id))
	if err.Error() == "user not found" {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

func (api *UserAPI) CreateUser(c echo.Context) error {
	var newUser model.User
	err := c.Bind(&newUser)
	fmt.Printf("request body: %+v", c.Request().Body)
	//err := json.NewDecoder(c.Request().Body).Decode(&newUser)
	if err != nil {
		fmt.Printf("\n%v\n", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request payload"})
	}
	user, err := api.userService.CreateUser(&newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}

func (api *UserAPI) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	var updatedUser model.User
	err = json.NewDecoder(c.Request().Body).Decode(&updatedUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request payload"})
	}
	user, err := api.userService.UpdateUser(int64(id), &updatedUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}
	return c.JSON(http.StatusOK, user)
}

func (api *UserAPI) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	err = api.userService.DeleteUser(int64(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
