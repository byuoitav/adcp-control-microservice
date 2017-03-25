package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func SetVolume(context echo.Context) error {
	return context.JSON(http.StatusOK, "")
}

func VolumeUnMute(context echo.Context) error {

	return context.JSON(http.StatusOK, "")
}
