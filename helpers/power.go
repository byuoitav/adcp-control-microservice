package helpers

import (
	"log"
	"strings"
)

func PowerOn(address string) error {

	log.Printf("Setting power of %v to on", address)
	command := "power \"on\""

	return sendCommand(command, address)
}

func PowerStandby(address string) error {
	log.Printf("Seting power of %v to off", address)
	command := "power \"off\""

	return sendCommand(command, address)
}

type Power struct {
	Power string `json:"status"`
}

func GetPowerStatus(address string) (Power, error) {

	log.Printf("Querying power state of %v", address)

	response, err := queryState("power_status ?", address)
	if err != nil {
		return Power{}, err
	}

	var status Power
	responseString := string(response)

	if strings.Contains(responseString, "standby") {
		status.Power = "standby"
	} else if strings.Contains(responseString, "startup") {
		status.Power = "startup"
	} else if strings.Contains(responseString, "on") {
		status.Power = "on"
	} else if strings.Contains(responseString, "cooling1") {
		status.Power = "cooling1"
	} else if strings.Contains(responseString, "cooling2") {
		status.Power = "cooling2"
	} else if strings.Contains(responseString, "saving_standby") {
		status.Power = "saving_standby"
	}

	return status, nil
}
