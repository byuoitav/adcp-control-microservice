package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/byuoitav/adcp-control-microservice/helpers"
	"github.com/byuoitav/common/log"
	se "github.com/byuoitav/common/status"
	"github.com/labstack/echo"
)

func setVolume(context echo.Context, pooled bool) error {
	address := context.Param("address")
	volumeLevel := context.Param("level")

	level, er := strconv.Atoi(volumeLevel)
	if er != nil {
		log.L.Warnf("Invalid volume level, non integer passed in request")
		return context.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid volume level %s. Must be in range 0-100. %s", volumeLevel, er.Error()))
	}

	err := helpers.SetVolume(address, level, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, se.Volume{level})
}

func powerOn(context echo.Context, pooled bool) error {
	address := context.Param("address")

	err := helpers.PowerOn(address, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, se.Power{"on"})

}

func powerStandby(context echo.Context, pooled bool) error {
	address := context.Param("address")

	err := helpers.PowerStandby(address, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, se.Power{"standby"})

}

func mute(context echo.Context, pooled bool) error {
	log.L.Debugf("Muting..")

	address := context.Param("address")

	err := helpers.SetMute(address, true, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, se.Mute{true})
}

func unMute(context echo.Context, pooled bool) error {
	log.L.Debugf("UnMuting..")

	address := context.Param("address")

	err := helpers.SetMute(address, false, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, se.Mute{false})
}

func displayBlank(context echo.Context, pooled bool) error {
	log.L.Debugf("Blanking Display..")

	address := context.Param("address")

	err := helpers.SetBlank(address, true, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, se.Blanked{true})
}

func displayUnBlank(context echo.Context, pooled bool) error {
	log.L.Debugf("Unblanking Display..")

	address := context.Param("address")

	err := helpers.SetBlank(address, false, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, se.Blanked{false})
}

func setInputPort(context echo.Context, pooled bool) error {
	log.L.Debugf("Setting input...")

	port := context.Param("port")
	address := context.Param("address")

	err := helpers.SetInput(address, port, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, se.Input{port})
}

func hasActiveSignal(context echo.Context, pooled bool) error {
	address := context.Param("address")
	port := context.Param("port")

	active, err := helpers.HasActiveSignal(address, port, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, active)
}

func getHardwareInfo(context echo.Context, pooled bool) error {
	address := context.Param("address")

	hardware, err := helpers.GetHardwareInfo(address, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, hardware)
}
