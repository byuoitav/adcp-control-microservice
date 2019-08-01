package adcp

import (
	"fmt"
	"time"

	"github.com/byuoitav/common/pooled"
	"github.com/byuoitav/common/status"
)

// SetPower .
func SetPower(address string, state status.Power) error {
	switch state.Power {
	case "standby":
		state.Power = "off"
	}

	work := func(conn pooled.Conn) error {
		conn.Log().Infof("Setting power state to %v", state.Power)

		cmd := []byte(fmt.Sprintf("power \"%s\"\r\n", state.Power))
		resp, err := writeAndRead(conn, cmd, 5*time.Second)
		if err != nil {
			return err
		}

		if resp != "ok" {
			return fmt.Errorf("unable to set power state to %v: %s", state.Power, resp)
		}

		conn.Log().Infof("Set power state to %v", state.Power)

		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return err
	}

	return nil
}

// SetBlanked .
func SetBlanked(address string, state status.Blanked) error {
	var str string
	switch state.Blanked {
	case true:
		str = "on"
	case false:
		str = "off"
	}

	work := func(conn pooled.Conn) error {
		conn.Log().Infof("Setting blanked state to %v", state.Blanked)

		cmd := []byte(fmt.Sprintf("blank \"%s\"\r\n", str))
		resp, err := writeAndRead(conn, cmd, 5*time.Second)
		if err != nil {
			return err
		}

		if resp != "ok" {
			return fmt.Errorf("unable to set blanked state to %v: %s", state.Blanked, resp)
		}

		conn.Log().Infof("Set blanked state to %v", state.Blanked)

		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return err
	}

	return nil
}

// SetInput .
func SetInput(address string, state status.Input) error {
	work := func(conn pooled.Conn) error {
		conn.Log().Infof("Setting input %v", state.Input)

		cmd := []byte(fmt.Sprintf("input \"%s\"\r\n", state.Input))
		resp, err := writeAndRead(conn, cmd, 5*time.Second)
		if err != nil {
			return err
		}

		if resp != "ok" {
			return fmt.Errorf("unable to set input to %v: %s", state.Input, resp)
		}

		conn.Log().Infof("Set input to %v", state.Input)

		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return err
	}

	return nil
}

// SetMuted .
func SetMuted(address string, state status.Mute) error {
	var str string
	switch state.Muted {
	case true:
		str = "on"
	case false:
		str = "off"
	}

	work := func(conn pooled.Conn) error {
		conn.Log().Infof("Setting muted state to %v", state.Muted)

		cmd := []byte(fmt.Sprintf("muting \"%s\"\r\n", str))
		resp, err := writeAndRead(conn, cmd, 5*time.Second)
		if err != nil {
			return err
		}

		if resp != "ok" {
			return fmt.Errorf("unable to set muted state to %v: %s", state.Muted, resp)
		}

		conn.Log().Infof("Set mute state to %v", state.Muted)

		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return err
	}

	return nil
}

// SetVolume .
// TODO some convert func?
func SetVolume(address string, volume status.Volume) error {
	work := func(conn pooled.Conn) error {
		conn.Log().Infof("Setting volume to %v", volume.Volume)

		cmd := []byte(fmt.Sprintf("volume %v\r\n", volume.Volume))
		resp, err := writeAndRead(conn, cmd, 5*time.Second)
		if err != nil {
			return err
		}

		if resp != "ok" {
			return fmt.Errorf("unable to set volume to %v: %s", volume.Volume, resp)
		}

		conn.Log().Infof("Set volume to %v", volume.Volume)

		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return err
	}

	return nil
}
