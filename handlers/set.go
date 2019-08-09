package handlers

import (
	"net/http"
	"strconv"

	"github.com/byuoitav/adcp-control-microservice/adcp"
	"github.com/byuoitav/common/status"
	"github.com/labstack/echo"
)

// SetPower .
func SetPower(ectx echo.Context) error {
	address := ectx.Param("address")
	state := status.Power{
		Power: ectx.Param("state"),
	}

	err := adcp.SetPower(address, state)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	return ectx.JSON(http.StatusOK, state)
}

// SetBlanked .
func SetBlanked(ectx echo.Context) error {
	address := ectx.Param("address")
	state := status.Blanked{
		Blanked: ectx.Param("state") == "blank",
	}

	err := adcp.SetBlanked(address, state)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	return ectx.JSON(http.StatusOK, state)
}

// SetInput .
func SetInput(ectx echo.Context) error {
	address := ectx.Param("address")
	state := status.Input{
		Input: ectx.Param("port"),
	}

	err := adcp.SetInput(address, state)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	return ectx.JSON(http.StatusOK, state)
}

// SetMuted .
func SetMuted(ectx echo.Context) error {
	address := ectx.Param("address")
	state := status.Mute{
		Muted: ectx.Param("state") == "mute",
	}

	err := adcp.SetMuted(address, state)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	return ectx.JSON(http.StatusOK, state)
}

// SetVolume .
func SetVolume(ectx echo.Context) error {
	address := ectx.Param("address")

	vol, err := strconv.Atoi(ectx.Param("level"))
	if err != nil {
		return ectx.String(http.StatusBadRequest, err.Error())
	}

	volume := status.Volume{
		Volume: vol,
	}

	err = adcp.SetVolume(address, volume)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	return ectx.JSON(http.StatusOK, volume)
}
