package helpers

import (
	"errors"
	"fmt"
	"log"
)

func SetVolume(address string, volumeLevel int) error {
	log.Printf("Setting volume of %s to %v", address, volumeLevel)

	if volumeLevel > 100 || volumeLevel < 0 {
		err := errors.New(fmt.Sprintf("Invalid volume level %v: must be in range 0-100", volumeLevel))
		log.Printf(err.Error())

		return err
	}
	command := fmt.Sprintf("volume %v", volumeLevel)

	return sendCommand(command, address)
}
