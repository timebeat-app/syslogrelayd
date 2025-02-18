package syslog_client

import (
	"fmt"
	"log"
	"log/syslog"
)

type Client struct {
	done       chan struct{}
	sysLoggers []*syslog.Writer
}

func NewSyslogClient(done chan struct{}, syslogServerConfig *SyslogServerConfig) *Client {

	controller := &Client{
		done: done,
	}

	for _, port := range syslogServerConfig.SyslogServerPorts {
		syslogServerConfig.SyslogServer.Port = int(port)
		sysLogger, err := syslog.Dial(syslogServerConfig.SyslogServer.Network(),
			syslogServerConfig.SyslogServer.AddrPort().String(),
			syslog.Priority(syslogServerConfig.SyslogAlertLevel),
			syslogServerConfig.SyslogTag)
		if err != nil {
			log.Fatalf("Unable to connect to syslog server: %s\n", err)
		}

		controller.sysLoggers = append(controller.sysLoggers, sysLogger)

	}
	return controller
}

func (controller *Client) Log(syslogMessage string) {

	for _, sysLogger := range controller.sysLoggers {
		if _, err := fmt.Fprintf(sysLogger, syslogMessage); err != nil {
			fmt.Printf("Syslog error: %s\n", err.Error())
		}
	}
}
