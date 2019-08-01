package handlers

import (
	"net/http"

	"github.com/byuoitav/adcp-control-microservice/adcp"
	"github.com/byuoitav/common/status"
	"github.com/labstack/echo"
)

// GetPower .
func GetPower(ectx echo.Context) error {
	address := ectx.Param("address")

	resp, err := adcp.GetPower(address)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	return ectx.JSON(http.StatusOK, resp)
}

// GetBlanked .
func GetBlanked(ectx echo.Context) error {
	address := ectx.Param("address")

	resp, err := adcp.GetBlanked(address)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	return ectx.JSON(http.StatusOK, resp)
}

// GetInput .
func GetInput(ectx echo.Context) error {
	address := ectx.Param("address")

	resp, err := adcp.GetInput(address)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	return ectx.JSON(http.StatusOK, resp)
}

// GetMuted .
func GetMuted(ectx echo.Context) error {
	address := ectx.Param("address")

	resp, err := adcp.GetMuted(address)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	return ectx.JSON(http.StatusOK, resp)
}

// GetVolume .
func GetVolume(ectx echo.Context) error {
	address := ectx.Param("address")

	resp, err := adcp.GetVolume(address)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	return ectx.JSON(http.StatusOK, resp)
}

// GetHardwareInfo .
func GetHardwareInfo(ectx echo.Context) error {
	address := ectx.Param("address")

	resp, err := adcp.GetHardwareInfo(address)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	return ectx.JSON(http.StatusOK, resp)
}

// GetActiveSignal .
func GetActiveSignal(ectx echo.Context) error {
	address := ectx.Param("address")

	resp, err := adcp.GetActiveSignal(address)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	return ectx.JSON(http.StatusOK, resp)
}
