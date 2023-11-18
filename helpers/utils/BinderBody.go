package utils

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

func BindBody[T any](c echo.Context, data *T) error {
	body := c.Request().Body
	decode := json.NewDecoder(body)
	return decode.Decode(&data)
}
