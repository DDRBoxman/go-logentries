package logentries

import (
	"fmt"
	"net"
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

func (l *Logentries) Write(p []byte) (n int, err error) {

	log := append([]byte(l.token+" "), p...)
	l.logs <- log

	return
}

func (l *Logentries) Close() {
	close(l.logs)
	<-l.done
	l.connection.Close()
}

func (l *Logentries) sendMessages() {
	for {
		log, more := <-l.logs
		if more {
			l.connection.Write(log)
		} else {
			break
		}
	}

	l.done <- true
}
