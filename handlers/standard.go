package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/byuoitav/adcp-control-microservice/helpers"
	se "github.com/byuoitav/av-api/statusevaluators"
	"github.com/byuoitav/common/log"
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

	return context.JSON(http.StatusOK, se.PowerStatus{"on"})

}

func powerStandby(context echo.Context, pooled bool) error {
	address := context.Param("address")

	err := helpers.PowerStandby(address, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, se.PowerStatus{"standby"})

}

func mute(context echo.Context, pooled bool) error {
	log.L.Debugf("Muting..")

	address := context.Param("address")

	err := helpers.SetMute(address, true, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, se.MuteStatus{true})
}

func unMute(context echo.Context, pooled bool) error {
	log.L.Debugf("UnMuting..")

	address := context.Param("address")

	err := helpers.SetMute(address, false, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, se.MuteStatus{false})
}

func displayBlank(context echo.Context, pooled bool) error {
	log.L.Debugf("Blanking Display..")

	address := context.Param("address")

	err := helpers.SetBlank(address, true, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, se.BlankedStatus{true})
}

func displayUnBlank(context echo.Context, pooled bool) error {
	log.L.Debugf("Unblanking Display..")

	address := context.Param("address")

	err := helpers.SetBlank(address, false, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, se.BlankedStatus{false})
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

func volumeLevel(context echo.Context, pooled bool) error {

	address := context.Param("address")

	level, err := helpers.GetVolumeLevel(address, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, level)
}

func muteStatus(context echo.Context, pooled bool) error {

	address := context.Param("address")

	status, err := helpers.GetMuteStatus(address, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status)
}

func powerStatus(context echo.Context, pooled bool) error {

	address := context.Param("address")

	status, err := helpers.GetPowerStatus(address, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status)
}

func blankedStatus(context echo.Context, pooled bool) error {
	address := context.Param("address")

	status, err := helpers.GetBlankStatus(address, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status)
}

func inputStatus(context echo.Context, pooled bool) error {
	address := context.Param("address")

	status, err := helpers.GetInputStatus(address, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status)
}
