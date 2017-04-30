package helpers

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Volume struct {
	Volume int `json:"volume"`
}

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

func GetVolumeLevel(address string) (Volume, error) {

	log.Printf("Querying volume of %s", address)

	resp, err := queryState("volume ?", address)
	if err != nil {
		return Volume{}, err
	}

	response := string(resp)
	fields := strings.Fields(response)
	level, err := strconv.Atoi(fields[0])
	if err != nil {
		return Volume{}, err
	}

	return Volume{Volume: level}, nil
}
