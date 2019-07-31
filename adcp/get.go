package adcp

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/byuoitav/common/pooled"
	"github.com/byuoitav/common/status"
)

// GetPower .
func GetPower(address string) (status.Power, error) {
	var status status.Power

	work := func(conn pooled.Conn) error {
		conn.Log().Infof("Getting power status")

		cmd := []byte(fmt.Sprintf("power_status ?\n"))
		resp, err := writeAndRead(conn, cmd, 5*time.Second)
		if err != nil {
			return err
		}

		conn.Log().Infof("Power status is %s", resp)

		// this list comes from adcp documentation
		switch resp {
		case `"standby"`:
			status.Power = "standby"
		case `"startup"`:
			status.Power = "on"
		case `"on"`:
			status.Power = "on"
		case `"cooling1"`:
			status.Power = "standby"
		case `"cooling2"`:
			status.Power = "standby"
		case `"saving_cooling1"`:
			status.Power = "standby"
		case `"saving_cooling2"`:
			status.Power = "standby"
		case `"saving_standby"`:
			status.Power = "standby"
		default:
			return fmt.Errorf("unknown power status '%s'", resp)
		}

		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return status, err
	}

	return status, nil
}

// GetBlanked .
func GetBlanked(address string) (status.Blanked, error) {
	var status status.Blanked

	work := func(conn pooled.Conn) error {
		conn.Log().Infof("Getting blanked status")

		cmd := []byte(fmt.Sprintf("blank ?\n"))
		resp, err := writeAndRead(conn, cmd, 5*time.Second)
		if err != nil {
			return err
		}

		conn.Log().Infof("Blanked status is %s", resp)

		switch resp {
		case `"on"`:
			status.Blanked = true
		case `"off"`:
			status.Blanked = false
		default:
			return fmt.Errorf("unknown blanked status '%s'", resp)
		}

		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return status, err
	}

	return status, nil
}

// GetInput .
func GetInput(address string) (status.Input, error) {
	var input status.Input

	work := func(conn pooled.Conn) error {
		conn.Log().Infof("Getting current input")

		cmd := []byte(fmt.Sprintf("blank ?\n"))
		resp, err := writeAndRead(conn, cmd, 5*time.Second)
		if err != nil {
			return err
		}

		conn.Log().Infof("Current input is %s", resp)

		input.Input = strings.Trim(resp, "\"")
		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return input, err
	}

	return input, nil
}

// GetMuted .
func GetMuted(address string) (status.Mute, error) {
	var status status.Mute

	work := func(conn pooled.Conn) error {
		conn.Log().Infof("Getting muted status")

		cmd := []byte(fmt.Sprintf("muting ?\n"))
		resp, err := writeAndRead(conn, cmd, 5*time.Second)
		if err != nil {
			return err
		}

		conn.Log().Infof("Muted status is %s", resp)

		switch resp {
		case `"on"`:
			status.Muted = true
		case `"off"`:
			status.Muted = false
		default:
			return fmt.Errorf("unknown muted status '%s'", resp)
		}

		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return status, err
	}

	return status, nil
}

// GetVolume .
func GetVolume(address string) (status.Volume, error) {
	var volume status.Volume

	work := func(conn pooled.Conn) error {
		conn.Log().Infof("Getting current volume")

		cmd := []byte(fmt.Sprintf("volume ?\n"))
		resp, err := writeAndRead(conn, cmd, 5*time.Second)
		if err != nil {
			return err
		}

		vol, err := strconv.Atoi(resp)
		if err != nil {
			return err
		}
		// TODO some convert function?

		conn.Log().Infof("Current volume is %d", vol)

		volume.Volume = vol
		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return volume, err
	}

	return volume, nil
}
