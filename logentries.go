package logentries

import (
	"net"
	"fmt"
)

type Logentries struct {
	token string
	port int
	ssl bool
	server string
}

const logentriesServer = "data.logentries.com"
const logentriesSecureServer = "api.logentries.com"

func New(token string) (l *Logentries) {
	l = new(Logentries)
	l.token = token
	l.server = logentriesServer
	l.port = 10000
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

func (l *Logentries) Write(p []byte) (n int, err error) {

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", l.server, l.port))
	if err != nil {
		fmt.Print(err)
	}
	conn.Write([]byte(l.token))
	conn.Write(p)
	return
}
