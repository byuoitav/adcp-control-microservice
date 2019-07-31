package helpers

import (
	"fmt"

	"github.com/byuoitav/common/log"

	"github.com/byuoitav/common/nerr"
)

type Mute struct {
	Mute string `json:"mute"`
}

func SetVolume(address string, volumeLevel int, pooled bool) *nerr.E {
	log.L.Infof("Setting volume of %s to %v", address, volumeLevel)

	if volumeLevel > 100 || volumeLevel < 0 {
		return nerr.Create(fmt.Sprintf("Invalid volume level %v: must be in range 0-100", volumeLevel), "params")
	}
	command := fmt.Sprintf("volume %v", volumeLevel)

	return sendCommand(command, address, pooled)
}

func SetMute(address string, muted bool, pooled bool) *nerr.E {
	var command string
	if muted {
		log.L.Infof("Muting %s", address)
		command = "muting \"on\""
	} else {
		log.L.Infof("Un-muting %s", address)
		command = "muting \"off\""
	}

	err := sendCommand(command, address, pooled)
	if err != nil {
		return err
	}

	return nil
}
