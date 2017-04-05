package helpers

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func getConnection(address string) (*net.TCPConn, error) {
	radder, err := net.ResolveTCPAddr("tcp", address+":53595")
	if err != nil {
		err = errors.New(fmt.Sprintf("error resolving address : %s", err.Error()))
		log.Printf(err.Error())
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, radder)
	if err != nil {

		err = errors.New(fmt.Sprintf("error dialing address : %s", err.Error()))
		log.Printf(err.Error())
		return nil, err
	}
	return conn, nil
}

func sendCommand(command string, address string) error {
	log.Printf("Sending Command %s", command)

	command = command + "\r\n"

	conn, err := getConnection(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = readUntil('\n', conn, 3)
	if err != nil {
		err = errors.New(fmt.Sprintf("Error reading first response on connect %s", err.Error()))
		log.Printf("%s", err.Error())
		return err
	}

	_, err = conn.Write([]byte(command))
	if err != nil {
		err = errors.New(fmt.Sprintf("Error sending command : %s", err.Error()))
		log.Printf("%s", err.Error())
		return err
	}
	resp, err := readUntil('\n', conn, 10)
	if err != nil {
		return err
	}

	if strings.Contains(string(resp), "ok") {
		log.Printf("Command Acknowledged")
		return nil
	}

	err = errors.New(fmt.Sprintf("Invalid response recieved: %s", resp))
	log.Printf("%s", err.Error())
	return err
}

func readUntil(delimeter byte, conn *net.TCPConn, timeoutInSeconds int) ([]byte, error) {

	conn.SetReadDeadline(time.Now().Add(time.Duration(int64(timeoutInSeconds)) * time.Second))

	buffer := make([]byte, 128)
	message := []byte{}

	for !charInBuffer(delimeter, buffer) {
		_, err := conn.Read(buffer)
		if err != nil {
			err = errors.New(fmt.Sprintf("Error reading response: %s", err.Error()))
			log.Printf("%s", err.Error())
			return message, err
		}

		message = append(message, buffer...)
	}
	return message, nil
}

func charInBuffer(toCheck byte, buffer []byte) bool {
	for _, b := range buffer {
		if toCheck == b {
			return true
		}
	}

	return false
}