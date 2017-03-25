package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/byuoitav/adcp-control-microservice/helpers"
	"github.com/labstack/echo"
)

func SetVolume(context echo.Context) error {

	address := context.Param("address")
	volumeLevel := context.Param("level")

	level, err := strconv.Atoi(volumeLevel)
	if err != nil {
		return context.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid volume level %s. Must be in range 0-100. %s", volumeLevel, err.Error()))
	}

	err = helpers.SetVolume(address, level)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, "ok")
}
