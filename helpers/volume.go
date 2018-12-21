package helpers

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/byuoitav/common/log"

	"github.com/byuoitav/common/nerr"
	se "github.com/byuoitav/common/status"
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

func GetVolumeLevel(address string, pooled bool) (se.Volume, *nerr.E) {

	log.L.Infof("Querying volume of %s", address)

	resp, err := queryState("volume ?", address, pooled)
	if err != nil {
		return se.Volume{}, err.Addf("Coudldn't get the volume level for %v", address)
	}

	response := string(resp)
	fields := strings.Fields(response)
	level, er := strconv.Atoi(fields[0])
	if er != nil {
		return se.Volume{}, nerr.Translate(er).Addf("Couldn't translate response %v to a volume level", fields[0])
	}

	return se.Volume{Volume: level}, nil
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

func GetMute(address string, pooled bool) (se.Mute, *nerr.E) {

	log.L.Infof("Querying mute status of %s", address)

	resp, err := queryState("muting ?", address, pooled)
	if err != nil {
		return se.Mute{}, err
	}

	response := string(resp)
	fields := strings.Fields(response)
	reg := regexp.MustCompile(`"([^"]*)"`)
	res := reg.ReplaceAllString(fields[0], "${1}")
	if res == "true" {
		return se.Mute{Muted: true}, nil
	} else {
		return se.Mute{Muted: false}, nil
	}
}
