package syslog_client

import "net"

type SyslogServerConfig struct {
	SyslogServer     *net.UDPAddr
	SyslogAlertLevel int
	SyslogTag        string
}
