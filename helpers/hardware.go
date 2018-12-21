package helpers

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/byuoitav/common/nerr"
	"github.com/byuoitav/common/structs"
)

// GetHardwareInfo gets the necessary hardware information from this device and returns it.
func GetHardwareInfo(address string, pooled bool) (structs.HardwareInfo, *nerr.E) {
	var toReturn structs.HardwareInfo

	// get the hostname
	addr, e := net.LookupAddr(address)
	if e != nil {
		toReturn.Hostname = address
	} else {
		toReturn.Hostname = strings.Trim(addr[0], ".")
	}

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

	toReturn.NetworkInfo.IPAddress = toString(ipBytes)

	// get the MAC address
	macBytes, err := queryState("mac_address ?", address, true)
	if err != nil {
		return toReturn, err.Add("Couldn't query the MAC address")
	}

	toReturn.NetworkInfo.MACAddress = toString(macBytes)

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

	e = json.Unmarshal(firmwareBytes, &firmware)
	if e != nil {
		return toReturn, nerr.Translate(e)
	}

	var firmwareString strings.Builder
	for _, f := range firmware {
		for _, v := range f {
			firmwareString.WriteString(fmt.Sprintf("%s$", strings.TrimSpace(v)))
		}
	}
	toReturn.FirmwareVersion = strings.TrimSuffix(firmwareString.String(), "$")

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
		return toReturn, nerr.Translate(e)
	}

	// get the error status
	errBytes, err := queryState("error ?", address, true)
	if err != nil {
		return toReturn, err.Add("Couldn't query the error status")
	}

	var errors []string

	e = json.Unmarshal(errBytes, &errors)
	if e != nil {
		return toReturn, nerr.Translate(e)
	}

	noWarnings, noErrors := noWarningsOrErrors(warnings, errors)

	if !noWarnings {
		toReturn.WarningStatus = warnings
	}
	if !noErrors {
		toReturn.ErrorStatus = errors
	}

	// get the power status
	powerBytes, err := queryState("power_status ?", address, true)
	if err != nil {
		return toReturn, err.Add("Couldn't query the power status")
	}

	toReturn.PowerStatus = toString(powerBytes)

	// get the timer info
	timerBytes, err := queryState("timer ?", address, true)
	if err != nil {
		return toReturn, err.Add("Couldn't query the timer information")
	}

	var timers []map[string]int

	e = json.Unmarshal(timerBytes, &timers)
	if e != nil {
		return toReturn, nerr.Translate(e)
	}

	toReturn.TimerInfo = timers

	return toReturn, nil
}

func toString(b []byte) string {
	return strings.Trim(string(b), "\"")
}

func trim(s string) string {
	return strings.Trim(s, "\"")
}

func noWarningsOrErrors(warns []string, errs []string) (warnings bool, errors bool) {
	warnings = false
	errors = false
	if len(warns) == 1 {
		warnings = strings.EqualFold(warns[0], "no_warn")
	}

	if len(errs) == 1 {
		errors = strings.EqualFold(errs[0], "no_err")
	}

	return warnings, errors
}
