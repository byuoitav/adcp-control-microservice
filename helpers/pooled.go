package helpers

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/nerr"
)

const bufferSize = 1000

const ttl = 20
const pruneinterval = 10

var connmap map[string]*Connection
var connmapMu *sync.Mutex

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

func init() {
	connmap = make(map[string]*Connection)
	connmapMu = &sync.Mutex{}
	go startGardener()
}

func MakeRequest(req Request) Response {

	req.ResponseChannel = make(chan Response, 1)

	log.L.Debugf("Making pooled request against %v", req.Address)

	conn, err := getPooledConnection(req.Address)
	if err != nil {
		return Response{
			Err: err.Addf("Couldn't get the connection to make request %v against %v", req.Command, req.Address),
		}
	}

	log.L.Debugf("Sending request down channel. Channel len: %v", len(conn.InChannel))
	//we write the request to the channel
	conn.InChannel <- req

	//now we wait
	resp := <-req.ResponseChannel

	log.L.Debugf("Response back.")

	if resp.Err != nil {
		log.L.Debugf("Error in response: %v", resp.Err.Error())
	}

	return resp
}

func getPooledConnection(addr string) (*Connection, *nerr.E) {
	connmapMu.Lock()
	v, ok := connmap[addr]
	if ok {
		log.L.Debugf("Using saved connection for %v", addr)
		connmapMu.Unlock()

		return v, nil
	}
	connmapMu.Unlock()

	return StartConnection(addr)
}

func StartConnection(address string) (*Connection, *nerr.E) {
	con, err := getConnection(address)
	if err != nil {
		return nil, err.Addf("Cannot get connection to start the conncetion minder")
	}

	log.L.Debugf("Reading first newline on connect")

	_, err = readUntil('\n', con, 3)
	if err != nil {
		return nil, err.Addf(fmt.Sprintf("Error reading first response on connect %s", err.Error()), "protocol")
	}

	conn := &Connection{
		Conn:              con,
		Address:           address,
		InChannel:         make(chan Request, bufferSize),
		SeppukuChannel:    make(chan bool, 1),
		LastCommunication: time.Now(),
	}

	go StartMinder(conn)

	connmapMu.Lock()
	connmap[address] = conn
	connmapMu.Unlock()

	return conn, nil
}

func StartMinder(conn *Connection) {

	//close the connection when we get out
	defer conn.Conn.Close()

	log.L.Infof("Starting minder for %v", conn.Address)
	for {
		conn.LastCommunication = time.Now()
		select {
		case req := <-conn.InChannel:
			log.L.Debugf("Handling request for: %v", conn.Address)
			//we make the request
			handleReq(conn, req)
			continue
		case <-conn.SeppukuChannel:
			log.L.Debugf("Starting minder close for %v", conn.Address)
			//remove yourself from the conn map, close your channel, empty it, and then close
			connmapMu.Lock()
			delete(connmap, conn.Address)
			connmapMu.Unlock()
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
		log.L.Debugf("Handling a query request")
		v, err := QueryStateWithConn(req.Command, req.Address, conn.Conn)
		req.ResponseChannel <- Response{
			Body: v,
			Err:  err,
		}
	} else {
		log.L.Debugf("Handling a Set State command")
		err := SendCommandWithConn(req.Command, req.Address, conn.Conn)
		req.ResponseChannel <- Response{
			Err: err,
		}
	}
}

func startGardener() {
	log.L.Infof("Starting gardener. Pruning interval is %v and ttl is %v", pruneinterval, ttl)

	ticker := time.NewTicker(time.Duration(pruneinterval) * time.Second)
	for {
		//wait for the ticker
		<-ticker.C

		log.L.Debugf("Running gardener")
		connmapMu.Lock()
		for k, v := range connmap {
			if time.Since(v.LastCommunication) > (time.Duration(ttl) * time.Second) {
				log.L.Debugf("Pruning connection for %v. Last communication over the channel was %v", k, v.LastCommunication)

				// time to kill it
				v.SeppukuChannel <- true
			}
		}
		connmapMu.Unlock()
	}
}
