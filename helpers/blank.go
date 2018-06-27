package helpers

import (
	"fmt"
	"log"
	"strings"

	"github.com/byuoitav/av-api/statusevaluators"
)

func SetBlank(address string, blank bool) error {
	log.Printf("Setting blank on %s to %v", address, blank)

	var command string
	if blank {
		command = fmt.Sprintf("blank \"on\"")
	} else {
		command = fmt.Sprintf("blank \"off\"")
	}

	return sendCommand(command, address)
}

func GetBlankStatus(address string) (statusevaluators.BlankedStatus, error) {
	log.Printf("Querying blank status of %s", address)

	response, err := queryState("blank ?", address)
	if err != nil {
		return statusevaluators.BlankedStatus{}, err
	}

	var status statusevaluators.BlankedStatus
	resp := string(response)

	if strings.Contains(resp, "on") {
		status.Blanked = true
	} else if strings.Contains(resp, "off") {
		status.Blanked = false
	}

	return status, nil
}
