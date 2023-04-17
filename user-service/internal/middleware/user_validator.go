package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/mail"
	"use/internal/model"

	"github.com/labstack/echo/v4"
)

func UserValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user model.User
		err := c.Bind(&user)
		if err != nil {
			return err
		}
		messages := make(map[string]string, 0)

		// validate email here
		if user.Email == "" {
			messages["email"] = "email is a required field."
		}

		if _, err := mail.ParseAddress(user.Email); err != nil {
			messages["email"] = "email is invalid."
		}

		// it is required field
		if user.FirstName == "" {
			messages["first_name"] = "first_name is a required field."
		}

		// it is required field
		if user.LastName == "" {
			messages["last_name"] = "last_name is a required field."
		}

		// it is required field
		if user.UserStatus == "" {
			messages["user_status"] = "user_status is a required field."
		}

		// it is required field
		if user.UserName == "" {
			messages["user_name"] = "user_name is a required field."
		}

		// it is required field
		if user.Department == "" {
			messages["department"] = "department is a required field."
		}

		if len(messages) > 0 {
			c.JSON(echo.ErrUnprocessableEntity.Code, messages)
			return errors.New("invalid request payload")
		}
		bodyBytes, _ := json.Marshal(user)
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		return next(c)
	}
}
