package syslog_client

import (
	"fmt"
	"log"
	"log/syslog"
	"net"
)

type Controller struct {
	done      chan struct{}
	sysLogger *syslog.Writer
}

func NewSyslogClient(done chan struct{}, remoteSyslogHost *net.UDPAddr) *Controller {

	controller := &Controller{
		done: done,
	}

	sysLogger, err := syslog.Dial(remoteSyslogHost.Network(),
		remoteSyslogHost.AddrPort().String(),
		syslog.LOG_ALERT, "Timebeat")
	if err != nil {
		log.Fatal(err)
	}

	controller.sysLogger = sysLogger
	return controller
}

func (controller *Controller) Log() {

	if _, err := fmt.Fprintf(controller.sysLogger, "This is a daemon alert with Timebeat tag."); err != nil {
		fmt.Printf("Syslog error: %s\n", err.Error())
	}

	if err := controller.sysLogger.Emerg("And this is a daemon emergency with Timebeat tag."); err != nil {
		fmt.Printf("Syslog error: %s\n", err.Error())
	}

}
