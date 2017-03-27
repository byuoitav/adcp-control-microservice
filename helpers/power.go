package helpers

import "log"

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
