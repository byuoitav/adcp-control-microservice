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
	var state status.Power

	work := func(conn pooled.Conn) error {
		conn.Log().Infof("Getting power state")

		cmd := []byte("power_status ?\r\n")
		resp, err := writeAndRead(conn, cmd, 5*time.Second)
		if err != nil {
			return err
		}

		conn.Log().Infof("Power state is %s", resp)

		// this list comes from adcp documentation
		switch resp {
		case `"standby"`:
			state.Power = "standby"
		case `"startup"`:
			state.Power = "on"
		case `"on"`:
			state.Power = "on"
		case `"cooling1"`:
			state.Power = "standby"
		case `"cooling2"`:
			state.Power = "standby"
		case `"saving_cooling1"`:
			state.Power = "standby"
		case `"saving_cooling2"`:
			state.Power = "standby"
		case `"saving_standby"`:
			state.Power = "standby"
		default:
			return fmt.Errorf("unknown power state '%s'", resp)
		}

		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return state, err
	}

	return state, nil
}

// GetBlanked .
func GetBlanked(address string) (status.Blanked, error) {
	var state status.Blanked

	work := func(conn pooled.Conn) error {
		conn.Log().Infof("Getting blanked state")

		cmd := []byte("blank ?\r\n")
		resp, err := writeAndRead(conn, cmd, 5*time.Second)
		if err != nil {
			return err
		}

		conn.Log().Infof("Blanked state is %s", resp)

		switch resp {
		case `"on"`:
			state.Blanked = true
		case `"off"`:
			state.Blanked = false
		default:
			return fmt.Errorf("unknown blanked state '%s'", resp)
		}

		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return state, err
	}

	return state, nil
}

// GetInput .
func GetInput(address string) (status.Input, error) {
	var input status.Input

	work := func(conn pooled.Conn) error {
		conn.Log().Infof("Getting current input")

		cmd := []byte("input ?\r\n")
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
	var state status.Mute

	work := func(conn pooled.Conn) error {
		conn.Log().Infof("Getting muted state")

		cmd := []byte("muting ?\r\n")
		resp, err := writeAndRead(conn, cmd, 5*time.Second)
		if err != nil {
			return err
		}

		conn.Log().Infof("Muted state is %s", resp)

		switch resp {
		case `"on"`:
			state.Muted = true
		case `"off"`:
			state.Muted = false
		default:
			return fmt.Errorf("unknown muted state '%s'", resp)
		}

		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return state, err
	}

	return state, nil
}

// GetVolume .
// TODO some convert function?
func GetVolume(address string) (status.Volume, error) {
	var volume status.Volume

	work := func(conn pooled.Conn) error {
		conn.Log().Infof("Getting current volume")

		cmd := []byte("volume ?\r\n")
		resp, err := writeAndRead(conn, cmd, 5*time.Second)
		if err != nil {
			return err
		}

		vol, err := strconv.Atoi(resp)
		if err != nil {
			return err
		}

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
