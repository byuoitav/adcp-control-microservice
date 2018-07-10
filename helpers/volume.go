package helpers

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	se "github.com/byuoitav/av-api/statusevaluators"
	"github.com/byuoitav/common/nerr"
)

type Mute struct {
	Mute string `json:"mute"`
}

func SetVolume(address string, volumeLevel int, pooled bool) *nerr.E {
	log.Printf("Setting volume of %s to %v", address, volumeLevel)

	if volumeLevel > 100 || volumeLevel < 0 {
		return nerr.Create(fmt.Sprintf("Invalid volume level %v: must be in range 0-100", volumeLevel), "params")
	}
	command := fmt.Sprintf("volume %v", volumeLevel)

	return sendCommand(command, address, pooled)
}

func GetVolumeLevel(address string, pooled bool) (se.Volume, *nerr.E) {

	log.Printf("Querying volume of %s", address)

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
		log.Printf("Muting %s", address)
		command = "muting \"on\""
	} else {
		log.Printf("Un-muting %s", address)
		command = "muting \"off\""
	}

	err := sendCommand(command, address, pooled)
	if err != nil {
		return err
	}

	return nil
}

func GetMuteStatus(address string, pooled bool) (se.MuteStatus, *nerr.E) {

	log.Printf("Querying mute status of %s", address)

	resp, err := queryState("muting ?", address, pooled)
	if err != nil {
		return se.MuteStatus{}, err
	}

	response := string(resp)
	fields := strings.Fields(response)
	reg := regexp.MustCompile(`"([^"]*)"`)
	res := reg.ReplaceAllString(fields[0], "${1}")
	if res == "true" {
		return se.MuteStatus{Muted: true}, nil
	} else {
		return se.MuteStatus{Muted: false}, nil
	}
}
