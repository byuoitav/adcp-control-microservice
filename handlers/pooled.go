package handlers

import (
	"github.com/labstack/echo"
)

// SetVolumePooled sets the volume when pooled
func SetVolumePooled(context echo.Context) error {
	return setVolume(context, true)
}

// SetVolume sets the volume
func SetVolume(context echo.Context) error {
	return setVolume(context, false)
}

// PowerOnPooled turns the power on when pooled
func PowerOnPooled(context echo.Context) error {
	return powerOn(context, true)
}

// PowerOn turns the power on
func PowerOn(context echo.Context) error {
	return powerOn(context, false)
}

// PowerStandbyPooled sets the power state to standby when pooled
func PowerStandbyPooled(context echo.Context) error {
	return powerStandby(context, true)
}

// PowerStandby sets the power state to standby
func PowerStandby(context echo.Context) error {
	return powerStandby(context, false)
}

// MutePooled mutes the sound when pooled
func MutePooled(context echo.Context) error {
	return mute(context, true)
}

// Mute mutes the sound
func Mute(context echo.Context) error {
	return mute(context, false)
}

// UnMutePooled unmutes the sound when pooled
func UnMutePooled(context echo.Context) error {
	return unMute(context, true)
}

// UnMute unmutes the sound
func UnMute(context echo.Context) error {
	return unMute(context, false)
}

// DisplayBlankPooled blanks the display when pooled
func DisplayBlankPooled(context echo.Context) error {
	return displayBlank(context, true)
}

// DisplayBlank blanks the display
func DisplayBlank(context echo.Context) error {
	return displayBlank(context, false)
}

// DisplayUnBlankPooled unblanks the display when pooled
func DisplayUnBlankPooled(context echo.Context) error {
	return displayUnBlank(context, true)
}

// DisplayUnBlank unblanks the display
func DisplayUnBlank(context echo.Context) error {
	return displayUnBlank(context, false)
}

// SetInputPortPooled sets the input port when pooled
func SetInputPortPooled(context echo.Context) error {
	return setInputPort(context, true)
}

// SetInputPort sets the input port
func SetInputPort(context echo.Context) error {
	return setInputPort(context, false)
}

// VolumeLevelPooled gets the volume level when pooled
func VolumeLevelPooled(context echo.Context) error {
	return volumeLevel(context, true)
}

// VolumeLevel gets the volume level
func VolumeLevel(context echo.Context) error {
	return volumeLevel(context, false)
}

// MuteStatusPooled checks if the sound is muted when pooled
func MuteStatusPooled(context echo.Context) error {
	return muteStatus(context, true)
}

// MuteStatus checks if the sound is muted
func MuteStatus(context echo.Context) error {
	return muteStatus(context, false)
}

// PowerStatusPooled checks the power status when pooled
func PowerStatusPooled(context echo.Context) error {
	return powerStatus(context, true)
}

// PowerStatus checks the power status
func PowerStatus(context echo.Context) error {
	return powerStatus(context, false)
}

// BlankedStatusPooled checks if the display is blanked when pooled
func BlankedStatusPooled(context echo.Context) error {
	return blankedStatus(context, true)
}

// BlankedStatus checks if the display is blanked
func BlankedStatus(context echo.Context) error {
	return blankedStatus(context, false)
}

// InputStatusPooled checks what input the display is on when pooled
func InputStatusPooled(context echo.Context) error {
	return inputStatus(context, true)
}

// InputStatus checks what input the display is on
func InputStatus(context echo.Context) error {
	return inputStatus(context, false)
}

// HasActiveInputPooled checks to see if the display has an active input when pooled
func HasActiveInputPooled(context echo.Context) error {
	return hasActiveInput(context, true)
}

// HasActiveInput checks to see if the display has an active input
func HasActiveInput(context echo.Context) error {
	return hasActiveInput(context, false)
}

// SerialNumberPooled gets the serial number when pooled
func SerialNumberPooled(context echo.Context) error {
	return serialNumber(context, true)
}

// SerialNumber gets the serial number
func SerialNumber(context echo.Context) error {
	return serialNumber(context, false)
}

// ModelNamePooled gets the model name when pooled
func ModelNamePooled(context echo.Context) error {
	return modelName(context, true)
}

// ModelName gets the model name
func ModelName(context echo.Context) error {
	return modelName(context, false)
}

// MACAddressPooled gets the MAC address when pooled
func MACAddressPooled(context echo.Context) error {
	return macAddress(context, true)
}

// MACAddress gets the MAC address
func MACAddress(context echo.Context) error {
	return macAddress(context, false)
}
