package helpers

import (
	"log"
	"strings"

	"github.com/byuoitav/av-api/statusevaluators"
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

func GetPowerStatus(address string) (statusevaluators.PowerStatus, error) {

	log.Printf("Querying power state of %v", address)

	response, err := queryState("power_status ?", address)
	if err != nil {
		return statusevaluators.PowerStatus{}, err
	}

	var status statusevaluators.PowerStatus
	responseString := string(response)

	if strings.Contains(responseString, "standby") {
		status.Power = "standby"
	} else if strings.Contains(responseString, "startup") {
		status.Power = "on"
	} else if strings.Contains(responseString, "on") {
		status.Power = "on"
	} else if strings.Contains(responseString, "cooling1") {
		status.Power = "standby"
	} else if strings.Contains(responseString, "cooling2") {
		status.Power = "standby"
	} else if strings.Contains(responseString, "saving_standby") {
		status.Power = "standby"
	}

	return status, nil
}
