package logentries

import (
	"io"
)

type ServiceNameLogger struct {
	writer      io.Writer
	serviceName string
}

func NewServiceNameLogger(serviceName string, writer io.Writer) (l *ServiceNameLogger) {
	l = new(ServiceNameLogger)
	l.serviceName = serviceName
	l.writer = writer

	return
}

func (l *ServiceNameLogger) Write(p []byte) (n int, err error) {
	appendedLog := append([]byte(l.serviceName+" "), p...)
	l.writer.Write(appendedLog)

	return
}
