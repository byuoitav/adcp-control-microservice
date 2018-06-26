package helpers

import (
	"bytes"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/nerr"
)

func getConnection(address string) (*net.TCPConn, *nerr.E) {

	radder, err := net.ResolveTCPAddr("tcp", address+":53595")
	if err != nil {
		return nil, nerr.Translate(err).Addf("error resolving address %v:", address)
	}

	conn, err := net.DialTCP("tcp", nil, radder)
	if err != nil {
		return nil, nerr.Translate(err).Addf("error dialing address %v", address)
	}
	return conn, nil
}

func sendCommand(command string, address string, pooled bool) *nerr.E {
	//if pooled we defer to something else else
	if pooled {

	}

	log.L.Debugf("Sending command %s", command)

	command = command + "\r\n"

	conn, err := getConnection(address)
	if err != nil {
		return err.Addf("Couldn't issue command %v to %v", command, address)
	}
	defer conn.Close()
	return SendCommandWithConn(command, address, conn)

}

func SendCommandWithConn(command, address string, conn *net.TCPConn) *nerr.E {
	_, err := readUntil('\n', conn, 3)
	if err != nil {
		return err.Addf("Error reading first response on connect")
	}

	_, er := conn.Write([]byte(command))
	if er != nil {
		return nerr.Translate(er).Addf("Error sending command")
	}
	resp, err := readUntil('\n', conn, 10)
	if err != nil {
		return nerr.Translate(err)
	}

	if strings.Contains(string(resp), "ok") {
		log.L.Debugf("Command Acknowledged")
		return nil
	}

	return nerr.Create(fmt.Sprintf("Invalid response recieved: %s", resp), "protocol")
}

func queryState(command string, address string, pooled bool) ([]byte, *nerr.E) {
	//if pooled we do something else
	if pooled {

	}

	log.L.Debugf("Sending command %s", command)

	command = command + "\r\n"

	connection, err := getConnection(address)
	if err != nil {
		return []byte{}, nerr.Translate(err).Addf("Couldn't query state %v of %v", command, address)
	}
	defer connection.Close()
	return QueryStateWithConn(command, address, connection)

}

func QueryStateWithConn(command, address string, conn *net.TCPConn) ([]byte, *nerr.E) {
	_, err := readUntil('\n', conn, 3)
	if err != nil {
		return []byte{}, err.Addf(fmt.Sprintf("Error reading first response on connect %s", err.Error()), "protocol")
	}

	_, er := conn.Write([]byte(command))
	if er != nil {
		return []byte{}, nerr.Translate(err).Addf("Error sending command %s to %v", command, address)
	}

	resp, err := readUntil('\n', conn, 10)
	if err != nil {
		return []byte{}, err.Addf("Timed out on read for command %v", command)
	}

	// trim specific chars off the response
	resp = bytes.Trim(resp, "\u0000")
	resp = bytes.Trim(resp, "\n")
	resp = bytes.Trim(resp, "\r")

	return resp, nil

}

func readUntil(delimeter byte, conn *net.TCPConn, timeoutInSeconds int) ([]byte, *nerr.E) {

	conn.SetReadDeadline(time.Now().Add(time.Duration(int64(timeoutInSeconds)) * time.Second))

	buffer := make([]byte, 128)
	message := []byte{}

	for !charInBuffer(delimeter, buffer) {
		_, err := conn.Read(buffer)
		if err != nil {
			return message, nerr.Translate(err).Add("Couldn't read until delimeter")
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
