package helpers

import (
	"strings"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/nerr"
)

// GetSerialNumber returns the serial number of the projector.
func GetSerialNumber(address string, pooled bool) (string, *nerr.E) {
	log.L.Debugf("Checking for the serial number of %s", address)

	response, err := queryState("serialnum ?", address, pooled)
	if err != nil {
		return "n/a", err.Add("Couldn't get the serial number")
	}

	return strings.Trim(string(response), "\""), nil
}

// GetModelName returns the model name of the projector.
func GetModelName(address string, pooled bool) (string, *nerr.E) {
	log.L.Debugf("Checking for the model name of %s", address)

	response, err := queryState("modelname ?", address, pooled)
	if err != nil {
		return "n/a", err.Add("Couldn't get the model name")
	}

	return strings.Trim(string(response), "\""), nil
}

// GetMACAddress returns the MAC address of the projector.
func GetMACAddress(address string, pooled bool) (string, *nerr.E) {
	log.L.Debugf("Checking for the MAC address of %s", address)

	response, err := queryState("mac_address ?", address, pooled)
	if err != nil {
		return "n/a", err.Add("Couldn't get the MAC address")
	}

	return strings.Trim(string(response), "\""), nil
}
