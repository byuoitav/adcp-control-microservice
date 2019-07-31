package helpers

import (
	"github.com/byuoitav/common/log"

	"github.com/byuoitav/common/nerr"
)

func PowerOn(address string, pooled bool) *nerr.E {
	log.L.Infof("Setting power of %v to on", address)
	command := "power \"on\""

	return sendCommand(command, address, pooled)
}

func PowerStandby(address string, pooled bool) *nerr.E {
	log.L.Infof("Seting power of %v to off", address)
	command := "power \"off\""

	return sendCommand(command, address, pooled)
}
