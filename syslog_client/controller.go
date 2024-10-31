package syslog_client

import (
	"fmt"
	"log"
	"log/syslog"
)

type Controller struct {
	done      chan struct{}
	sysLogger *syslog.Writer
}

func NewSyslogClient(done chan struct{}, syslogServerConfig *SyslogServerConfig) *Controller {

	controller := &Controller{
		done: done,
	}

	sysLogger, err := syslog.Dial(syslogServerConfig.SyslogServer.Network(),
		syslogServerConfig.SyslogServer.AddrPort().String(),
		syslog.Priority(syslogServerConfig.SyslogAlertLevel),
		syslogServerConfig.SyslogTag)
	if err != nil {
		log.Fatalf("Unable to connect to syslog server: %s\n", err)
	}

	controller.sysLogger = sysLogger
	return controller
}

func (controller *Controller) Log(syslogMessage string) {

	if _, err := fmt.Fprintf(controller.sysLogger, syslogMessage); err != nil {
		fmt.Printf("Syslog error: %s\n", err.Error())
	}
}
