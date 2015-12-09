package logentries

import (
	"fmt"
	"net"
	"time"
)

type Logentries struct {
	token      string
	port       int
	ssl        bool
	server     string
	logs       chan []byte
	done       chan bool
	connection net.Conn
}

const logentriesServer = "data.logentries.com"
const logentriesSecureServer = "api.logentries.com"

// Create a new logentries instance
func New(token string) (l *Logentries) {
	l = new(Logentries)
	l.token = token
	l.server = logentriesServer
	l.port = 10000
	l.logs = make(chan []byte, 50)
	l.done = make(chan bool)

	l.connect()
	go l.sendMessages()

	return
}

// Set the port to send data to Logentries on
//
// Valid ports: 80, 514, 10000, 20000
//
// 20000 automatically enables SSL
func (l *Logentries) Port(port int) {
	if port == 20000 {
		l.ssl = true
		l.server = logentriesServer
	} else {
		l.ssl = false
		l.server = logentriesSecureServer
	}
	l.port = port
}

// Use SSL when sending data to Logentries
//
// Sets port to 20000
func (l *Logentries) UseSSL(useSSL bool) {
	l.ssl = useSSL
	l.port = 20000
	l.server = logentriesSecureServer
}

func (l *Logentries) connect() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", l.server, l.port))
	if err != nil {
		fmt.Print(err)
	}

	l.connection = conn
}

// Implement the io.Writer interface
func (l *Logentries) Write(p []byte) (n int, err error) {

	log := append([]byte(l.token+" "), p...)
	l.logs <- log

	return
}

// Clean up the logger and send any remaining messages
func (l *Logentries) Close() {
	close(l.logs)
	<-l.done
	l.connection.Close()
}

func (l *Logentries)ensureConnection() {
	buf := make([]byte, 1)

	l.connection.SetReadDeadline(time.Now())

	_, err := l.connection.Read(buf)
	switch err.(type) {
	case net.Error:
		if err.(net.Error).Timeout() == true {
			l.connection.SetReadDeadline(time.Time{})
			return
		}
	}

	l.connect()
}

func (l *Logentries) sendMessages() {
	for {
		log, more := <-l.logs
		if more {
			l.ensureConnection()
			l.connection.Write(log)
		} else {
			break
		}
	}

	l.done <- true
}
