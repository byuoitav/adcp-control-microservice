package helpers

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/byuoitav/av-api/statusevaluators"
)

var validADCPInputs = []string{
	"video1",
	"svideo1",
	"rgb1",
	"rgb2",
	"dvi1",
	"hdmi1",
	"hdmi2",
	"network",
	"usb_a",
	"usb_b",
	"hdbaset1",
	"option1",
}

func SetInput(address, port string) error {
	log.Printf("Setting input on %s to %s", address, port)

	validInput := false
	for _, input := range validADCPInputs {
		if strings.EqualFold(port, input) {
			validInput = true
			break
		}
	}

	if !validInput {
		return errors.New(fmt.Sprintf("error: %s is not a valid ADCP input.", port))
	}

	command := fmt.Sprintf("input \"%s\"", port)
	return sendCommand(command, address)
}

func GetInputStatus(address string) (statusevaluators.Input, error) {
	log.Printf("Querying input status of %s", address)

	response, err := queryState("input ?", address)
	if err != nil {
		return statusevaluators.Input{}, nil
	}

	status := statusevaluators.Input{
		Input: strings.Trim(string(response), "\""),
	}
	return status, nil
}
