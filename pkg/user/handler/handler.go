package handler

import (
	"fmt"
	"net/http"
	"reflect"
	"s3service/pkg/user/service"

	"github.com/labstack/echo"
)

type Handler struct {
	Service service.Service
}
type UserHandler interface {
	Upload(c echo.Context) error
	Retrive(c echo.Context) error
}

func (h *Handler) Upload(c echo.Context) error {
	file, _ := c.FormFile("file")
	fmt.Println("file  ....", file)
	fmt.Println("Type of x:", reflect.TypeOf(file))
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer src.Close()

	err = h.Service.Upload(src, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.String(http.StatusOK, "File uploaded successfully")

}
func (h *Handler) Retrive(c echo.Context) error {
	name := c.Param("name")
	url, err := h.Service.Retrive(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})

	}

	return c.String(http.StatusOK, url)

}
