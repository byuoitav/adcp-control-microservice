package adcp

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/byuoitav/common/pooled"
)

var pool = pooled.NewMap(100*time.Second, 300*time.Millisecond, getConnection)

const (
	// CR is a carriage return
	CR = '\r'
	// LF is a line feed
	LF = '\n'
)

func getConnection(key interface{}) (pooled.Conn, error) {
	address, ok := key.(string)
	if !ok {
		return nil, fmt.Errorf("key must be a string")
	}

	conn, err := net.DialTimeout("tcp", address+":53595", 10*time.Second)
	if err != nil {
		return nil, err
	}

	return pooled.Wrap(conn), nil
}

func writeAndRead(conn pooled.Conn, cmd []byte, readTimeout time.Duration) (string, error) {
	n, err := conn.Write(cmd)
	switch {
	case err != nil:
		return "", err
	case n != len(cmd):
		return "", fmt.Errorf("wrote %v/%v bytes of command 0x%x", n, len(cmd), cmd)
	}

	b, err := conn.ReadUntil(LF, readTimeout)
	if err != nil {
		return "", err
	}

	conn.Log().Debugf("Response from command: 0x%x", b)
	return strings.TrimSpace(string(b)), nil
}
