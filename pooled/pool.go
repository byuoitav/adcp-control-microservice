package pooled

import (
	"net"
	"time"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/nerr"
)

const bufferSize = 1000

const timout = 600

var connmap map[string]*Connection

type Connection struct {
	Conn              *net.TCPConn
	Address           string
	InChannel         chan Request
	SeppukuChannel    chan bool
	LastCommunication time.Time
}

type Request struct {
	Command         string
	Address         string
	Query           bool
	ResponseChannel chan Response
}

type Response struct {
	Body []byte
	Err  *nerr.E
}

func getConnection(addr string) (*net.TCPConn, *nerr.E) {

	//Build a new TCP connection with the address
	radder, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, nerr.Translate(err).Addf("Couldn't resolve address %v", addr)
	}

	conn, err := net.DialTCP("tcp", nil, radder)
	if err != nil {
		return nil, nerr.Translate(err).Addf("Couldn't dial address %v", addr)
	}

	return conn, nil
}

func StartConnection(address string) (*Connection, *nerr.E) {
	conn, err := getConnection(address)
	if err != nil {
		return nil, err.Addf("Cannot get connection to start the conncetion minder")
	}

	conn := &Connection{
		Conn:              conn,
		Address:           addr,
		InChannel:         make(chan Request, bufferSize),
		SeppukuChannel:    make(chan bool, 1),
		LastCommunication: time.Now(),
	}

	go StartMinder(conn)

	return conn
}

func StartMinder(conn *Connection) {
	log.L.Infof("Starting minder for %v", conn.Address)
	for {
		select {
		case req := <-conn.InChannel:
			log.L.Debugf("Handling request for: %v", conn.Address)
			//we make the request
			handleReq(conn, req)
			continue
		case <-conn.SeppukuChannel:
			log.L.Debugf("Starting minder close for %v", conn.Address)
			//remove yourself from the conn map, close your channel, empty it, and then close
			delete(connmap, conn.Address)
			close(conn.InChannel)
			for req := range conn.InChannel {
				log.L.Debugf("Clearing (handling) request for: %v", conn.Address)
				handleReq(conn, req)
			}

			log.L.Infof("Closing minder for %v", conn.Address)
			return
		}
	}
}

func handleReq(conn *Connection, req Request) {

	if req.Query {
		v, err := QueryStateWithConn(req.Command, req.Address, conn.Conn)
		req.ResponseChannel <- Response{
			Body: v,
			Err:  err,
		}
	} else {
		err := SendCommandWithCon(req.Command, req.Address, conn.Conn)
		req.ResponseChannel <- Response{
			Err: err,
		}
	}
}
