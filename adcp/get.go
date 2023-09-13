package adcp

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/byuoitav/common/pooled"
	"github.com/byuoitav/common/status"
	"github.com/byuoitav/common/structs"
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

	volume.Volume = adcpToNormalVolume(volume.Volume)
	return volume, nil
}

// GetHardwareInfo .
func GetHardwareInfo(address string) (structs.HardwareInfo, error) {
	var info structs.HardwareInfo

	work := func(conn pooled.Conn) error {
		conn.Log().Infof("Getting hardware info")

		// model name
		cmd := []byte("modelname ?\r\n")
		resp, err := writeAndRead(conn, cmd, 3*time.Second)
		if err != nil {
			return err
		}

		info.ModelName = strings.Trim(resp, "\"")

		// ip address
		cmd = []byte("ipv4_ip_address ?\r\n")
		resp, err = writeAndRead(conn, cmd, 3*time.Second)
		if err != nil {
			return err
		}

		info.NetworkInfo.IPAddress = strings.Trim(resp, "\"")

		// gateway
		cmd = []byte("ipv4_default_gateway ?\r\n")
		resp, err = writeAndRead(conn, cmd, 3*time.Second)
		if err != nil {
			return err
		}

		info.NetworkInfo.Gateway = strings.Trim(resp, "\"")

		// dns
		cmd = []byte("ipv4_dns_server1 ?\r\n")
		resp, err = writeAndRead(conn, cmd, 3*time.Second)
		if err != nil {
			return err
		}

		info.NetworkInfo.DNS = append(info.NetworkInfo.DNS, strings.Trim(resp, "\""))

		cmd = []byte("ipv4_dns_server2 ?\r\n")
		resp, err = writeAndRead(conn, cmd, 3*time.Second)
		if err != nil {
			return err
		}

		info.NetworkInfo.DNS = append(info.NetworkInfo.DNS, strings.Trim(resp, "\""))

		// mac address
		cmd = []byte("mac_address ?\r\n")
		resp, err = writeAndRead(conn, cmd, 3*time.Second)
		if err != nil {
			return err
		}

		info.NetworkInfo.MACAddress = strings.Trim(resp, "\"")

		// serial number
		cmd = []byte("serialnum ?\r\n")
		resp, err = writeAndRead(conn, cmd, 3*time.Second)
		if err != nil {
			return err
		}

		info.SerialNumber = strings.Trim(resp, "\"")

		// filter status
		cmd = []byte("filter_status ?\r\n")
		resp, err = writeAndRead(conn, cmd, 3*time.Second)
		if err != nil {
			return err
		}

		info.FilterStatus = strings.Trim(resp, "\"")

		// power status
		cmd = []byte("power_status ?\r\n")
		resp, err = writeAndRead(conn, cmd, 3*time.Second)
		if err != nil {
			return err
		}

		info.PowerStatus = strings.Trim(resp, "\"")

		// warnings
		cmd = []byte("warning ?\r\n")
		resp, err = writeAndRead(conn, cmd, 3*time.Second)
		if err != nil {
			return err
		}

		if resp[0] == '[' {
			err = json.Unmarshal([]byte(resp), &info.WarningStatus)
		}
		if err != nil {
			return err
		}

		// errors
		cmd = []byte("error ?\r\n")
		resp, err = writeAndRead(conn, cmd, 3*time.Second)
		if err != nil {
			return err
		}

		err = json.Unmarshal([]byte(resp), &info.ErrorStatus)
		if err != nil {
			return err
		}

		// timer info
		cmd = []byte("timer ?\r\n")
		resp, err = writeAndRead(conn, cmd, 3*time.Second)
		if err != nil {
			return err
		}

		err = json.Unmarshal([]byte(resp), &info.TimerInfo)
		if err != nil {
			return err
		}

		conn.Log().Infof("Hardware info %+v", info)

		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return info, err
	}

	return info, nil
}

// GetActiveSignal .
func GetActiveSignal(address string) (structs.ActiveSignal, error) {
	var state structs.ActiveSignal

	work := func(conn pooled.Conn) error {
		conn.Log().Infof("Getting active signal")

		cmd := []byte("signal ?\r\n")
		resp, err := writeAndRead(conn, cmd, 5*time.Second)
		if err != nil {
			return err
		}

		conn.Log().Infof("Active signal is %s", resp)

		switch resp {
		case `"Invalid"`:
			state.Active = false
		case "ok":
			state.Active = false
		default:
			state.Active = true
		}

		return nil
	}

	err := pool.Do(address, work)
	if err != nil {
		return state, err
	}

	return state, nil
}
