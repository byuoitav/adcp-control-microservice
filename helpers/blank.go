package helpers

import (
	"fmt"

	"github.com/byuoitav/common/log"

	"github.com/byuoitav/common/nerr"
)

func SetBlank(address string, blank, pooled bool) *nerr.E {
	log.L.Infof("Setting blank on %s to %v", address, blank)

	var command string
	if blank {
		command = fmt.Sprintf("blank \"on\"")
	} else {
		command = fmt.Sprintf("blank \"off\"")
	}

	return sendCommand(command, address, pooled)
}
