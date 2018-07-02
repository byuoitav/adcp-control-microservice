package helpers

import (
	"log"
	"strings"

	"github.com/byuoitav/av-api/statusevaluators"
	"github.com/byuoitav/common/nerr"
	"github.com/fatih/color"
)

func PowerOn(address string, pooled bool) *nerr.E {
	log.Printf("Setting power of %v to on", address)
	command := "power \"on\""

	return sendCommand(command, address, pooled)
}

func PowerStandby(address string, pooled bool) *nerr.E {
	log.Printf("Seting power of %v to off", address)
	command := "power \"off\""

	return sendCommand(command, address, pooled)
}

func GetPowerStatus(address string, pooled bool) (statusevaluators.PowerStatus, *nerr.E) {

	log.Printf("%s", color.HiCyanString("[helpers] querying power state of %v", address))

	response, err := queryState("power_status ?", address, pooled)
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
