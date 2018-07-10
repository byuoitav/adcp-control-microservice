package handlers

import (
	"github.com/labstack/echo"
)

func SetVolumePooled(context echo.Context) error {
	return setVolume(context, true)
}

func SetVolume(context echo.Context) error {
	return setVolume(context, false)
}

func PowerOnPooled(context echo.Context) error {
	return powerOn(context, true)
}

func PowerOn(context echo.Context) error {
	return powerOn(context, false)
}

func PowerStandbyPooled(context echo.Context) error {
	return powerStandby(context, true)
}

func PowerStandby(context echo.Context) error {
	return powerStandby(context, false)
}

func MutePooled(context echo.Context) error {
	return mute(context, true)
}

func Mute(context echo.Context) error {
	return mute(context, false)
}

func UnMutePooled(context echo.Context) error {
	return unMute(context, true)
}
func UnMute(context echo.Context) error {
	return unMute(context, false)
}

func DisplayBlankPooled(context echo.Context) error {
	return displayBlank(context, true)
}

func DisplayBlank(context echo.Context) error {
	return displayBlank(context, false)
}

func DisplayUnBlankPooled(context echo.Context) error {
	return displayUnBlank(context, true)
}

func DisplayUnBlank(context echo.Context) error {
	return displayUnBlank(context, false)
}

func SetInputPortPooled(context echo.Context) error {
	return setInputPort(context, true)
}

func SetInputPort(context echo.Context) error {
	return setInputPort(context, false)
}

func VolumeLevelPooled(context echo.Context) error {
	return volumeLevel(context, true)
}

func VolumeLevel(context echo.Context) error {
	return volumeLevel(context, false)
}

func MuteStatusPooled(context echo.Context) error {
	return muteStatus(context, true)
}
func MuteStatus(context echo.Context) error {
	return muteStatus(context, false)
}

func PowerStatusPooled(context echo.Context) error {
	return powerStatus(context, true)
}
func PowerStatus(context echo.Context) error {
	return powerStatus(context, false)
}

func BlankedStatusPooled(context echo.Context) error {
	return blankedStatus(context, true)
}
func BlankedStatus(context echo.Context) error {
	return blankedStatus(context, false)
}

func InputStatusPooled(context echo.Context) error {
	return inputStatus(context, true)
}

func InputStatus(context echo.Context) error {
	return inputStatus(context, false)
}
