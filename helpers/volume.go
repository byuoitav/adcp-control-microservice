package helpers

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"regexp"
)

type Volume struct {
	Volume int `json:"volume"`
}

type Mute struct {
	Mute string `json:"mute"`
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

func SetMute(address string, muted bool) error {

	var command string
	if muted {
		log.Printf("Muting %s", address)
		command = "muting \"on\""
	} else {
		log.Printf("Un-muting %s", address) 
		command = "muting \"off\""
	}

	err := sendCommand(command, address)
	if err != nil {
		 return err
	 }

	 return nil
 }

func GetMuteStatus(address string) (Mute, error){

	log.Printf("Querying mute status of %s", address)

	resp, err := queryState("muting ?", address)
	if err != nil {
		return Mute{}, err
	}

	response := string(resp)
	fields := strings.Fields(response)
	reg := regexp.MustCompile(`"([^"]*)"`)
    res := reg.ReplaceAllString(fields[0], "${1}")
	return Mute{Mute: res}, nil
}
