package syslog_client

import "net"

type SyslogServerConfig struct {
	SyslogServerPorts []uint16
	SyslogServer      *net.UDPAddr
	SyslogAlertLevel  int
	SyslogTag         string
}
