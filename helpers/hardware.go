package helpers

import (
	"encoding/json"
	"strings"

	"github.com/byuoitav/common/nerr"
	"github.com/byuoitav/common/structs"
)

// GetHardwareInfo gets the necessary hardware information from this device and returns it.
func GetHardwareInfo(address string, pooled bool) (structs.HardwareInfo, *nerr.E) {
	var toReturn structs.HardwareInfo

	// get the model name
	modelBytes, err := queryState("modelname ?", address, pooled)
	if err != nil {
		return toReturn, err.Add("Couldn't query the model name")
	}

	toReturn.ModelName = toString(modelBytes)

	// get the IP address
	ipBytes, err := queryState("ipv4_ip_address ?", address, true)
	if err != nil {
		return toReturn, err.Add("Couldn't query the IP address")
	}

	toReturn.IPAddress = toString(ipBytes)

	// get the MAC address
	macBytes, err := queryState("mac_address ?", address, true)
	if err != nil {
		return toReturn, err.Add("Couldn't query the MAC address")
	}

	toReturn.MACAddress = toString(macBytes)

	// get the serial number
	serialBytes, err := queryState("serialnum ?", address, true)
	if err != nil {
		return toReturn, err.Add("Couldn't query the serial number")
	}

	toReturn.SerialNumber = toString(serialBytes)

	// get the firmware version
	firmwareBytes, err := queryState("version ?", address, true)
	if err != nil {
		return toReturn, err.Add("Couldn't query the firmware version")
	}

	var firmware []map[string]string

	e := json.Unmarshal(firmwareBytes, &firmware)
	if e != nil {
		return toReturn, err.Add(e.Error())
	}

	toReturn.FirmwareVersion = firmware

	// get the filter status
	filterBytes, err := queryState("filter_status ?", address, true)
	if err != nil {
		return toReturn, err.Add("Couldn't query the filter status")
	}

	toReturn.FilterStatus = toString(filterBytes)

	// get the warning status
	warnBytes, err := queryState("warning ?", address, true)
	if err != nil {
		return toReturn, err.Add("Couldn't query the warning status")
	}

	var warnings []string

	e = json.Unmarshal(warnBytes, &warnings)
	if e != nil {
		return toReturn, err.Add(e.Error())
	}

	toReturn.WarningStatus = warnings

	// get the error status
	errBytes, err := queryState("error ?", address, true)
	if err != nil {
		return toReturn, err.Add("Couldn't query the error status")
	}

	var errors []string

	e = json.Unmarshal(errBytes, &errors)
	if e != nil {
		return toReturn, err.Add(e.Error())
	}

	toReturn.ErrorStatus = errors

	return toReturn, nil
}

func toString(b []byte) string {
	return strings.Trim(string(b), "\"")
}

func trim(s string) string {
	return strings.Trim(s, "\"")
}
