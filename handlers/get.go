package handlers

import (
	"net/http"

	"github.com/byuoitav/adcp-control-microservice/adcp"
	"github.com/labstack/echo"
)

// GetPower .
func GetPower(ectx echo.Context) error {
	address := ectx.Param("address")

	status, err := adcp.GetPower(address)
	if err != nil {
		return ectx.String(http.StatusInternalServerError, err.Error())
	}

	return ectx.JSON(http.StatusOK, status)
}

// GetBlanked .
func GetBlanked(ectx echo.Context) error {
	address := ectx.Param("address")

	status, err := adcp.GetBlanked(address)
	if err != nil {
		return ectx.String(http.StatusInternalServerError, err.Error())
	}

	return ectx.JSON(http.StatusOK, status)
}

// GetInput .
func GetInput(ectx echo.Context) error {
	address := ectx.Param("address")

	status, err := adcp.GetInput(address)
	if err != nil {
		return ectx.String(http.StatusInternalServerError, err.Error())
	}

	return ectx.JSON(http.StatusOK, status)
}

// GetMuted .
func GetMuted(ectx echo.Context) error {
	address := ectx.Param("address")

	status, err := adcp.GetMuted(address)
	if err != nil {
		return ectx.String(http.StatusInternalServerError, err.Error())
	}

	return ectx.JSON(http.StatusOK, status)
}

// GetVolume .
func GetVolume(ectx echo.Context) error {
	address := ectx.Param("address")

	status, err := adcp.GetVolume(address)
	if err != nil {
		return ectx.String(http.StatusInternalServerError, err.Error())
	}

	return ectx.JSON(http.StatusOK, status)
}
