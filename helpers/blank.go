package helpers

import (
	"fmt"
	"log"
	"strings"

	"github.com/byuoitav/av-api/statusevaluators"
	"github.com/byuoitav/common/nerr"
)

func SetBlank(address string, blank, pooled bool) *nerr.E {
	log.Printf("Setting blank on %s to %v", address, blank)

	var command string
	if blank {
		command = fmt.Sprintf("blank \"on\"")
	} else {
		command = fmt.Sprintf("blank \"off\"")
	}

	return sendCommand(command, address, pooled)
}

func GetBlankStatus(address string, pooled bool) (statusevaluators.BlankedStatus, *nerr.E) {
	log.Printf("Querying blank status of %s", address)

	response, err := queryState("blank ?", address, pooled)
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
