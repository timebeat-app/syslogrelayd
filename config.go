package main

import (
	"github.com/timebeat-app/syslogrelayd/http_server"
	"github.com/timebeat-app/syslogrelayd/syslog_client"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

// This go generate probably won't be used, but cool nonetheless
/*
//go:generate go run golang.org/x/tools/cmd/stringer -type=Priority

type Priority int

const (
	LOG_EMERG Priority = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)
*/

var appConfig AppConfig

type AppConfig struct {
	syslogServer syslog_client.SyslogServerConfig
	httpServer   http_server.HttpServerConfig
}

/*
 * Configurable via environment variables :
 * SYSLOGRELAYD_SYSLOG_HOST 		// Mandatory
 * SYSLOGRELAYD_SYSLOG_PORT 		// Default: 514
 * SYSLOGRELAYD_SYSLOG_ALERT_LEVEL	// Default: alert (1)
 * SYSLOGRELAYD_SYSLOG_TAG 			// Default: Timebeat
 * SYSLOGRELAYD_HTTP_PORT 			// Default: 8080
 * SYSLOGRELAYD_HTTP_URL_PATH 		// Default: /
 */

func parseConfig() {
	parseSyslogUDPAddr()
	parseSyslogAlertLevel()
	parseSyslogTag()
	parseHTTPPort()
	parseHTTPURLPath()
}

func parseHTTPURLPath() {
	httpURLPathEnv := os.Getenv("SYSLOGRELAYD_HTTP_URL_PATH")
	if httpURLPathEnv == "" {
		httpURLPathEnv = "/"
	}
	appConfig.httpServer.HTTPURLPath = httpURLPathEnv
}

func parseHTTPPort() {

	// Port
	httpPort := 8080
	httpPortEnv := os.Getenv("SYSLOGRELAYD_HTTP_PORT")
	if convertedPort, err := strconv.ParseInt(httpPortEnv, 10, 16); err == nil {
		httpPort = int(convertedPort)
	}
	appConfig.httpServer.HTTPServerPort = httpPort
}

func parseSyslogTag() {

	tag := os.Getenv("SYSLOGRELAYD_SYSLOG_TAG")
	if tag == "" {
		tag = "Timebeat"
	}
	appConfig.syslogServer.SyslogTag = tag
}

func parseSyslogAlertLevel() {

	syslogPriorityEnv := strings.ToLower(os.Getenv("SYSLOGRELAYD_SYSLOG_ALERT_LEVEL"))

	var syslogPriority int

	switch syslogPriorityEnv {

	case "emergency":
		syslogPriority = 0

	case "alert":
		syslogPriority = 1

	case "critical":
		syslogPriority = 2

	case "error":
		syslogPriority = 3

	case "warning":
		syslogPriority = 4

	case "notification":
		syslogPriority = 5

	case "informational":
		syslogPriority = 6

	case "debugging":
		syslogPriority = 7

	default:
		syslogPriority = 1
	}

	appConfig.syslogServer.SyslogAlertLevel = syslogPriority
}

func parseSyslogUDPAddr() {

	// Host
	syslogHostEnv := os.Getenv("SYSLOGRELAYD_SYSLOG_HOST")
	if syslogHostEnv == "" {
		log.Fatalf("Syslog host is not set. Pleae declare the " +
			"SYSLOGRELAYD_SYSLOG_HOST environment variable\n")
	}
	var syslogIP net.IP

	if addrs, err := net.LookupHost(syslogHostEnv); len(addrs) > 0 && err == nil {
		syslogIP = net.ParseIP(addrs[0])
	} else {
		syslogIP = net.ParseIP(syslogHostEnv)
	}

	if syslogIP == nil {
		log.Fatalf("Unable to resolve syslog host: %s\n", syslogHostEnv)
	}

	// Port
	syslogDefaultPort := uint16(514)
	var syslogPorts []uint16

	syslogPortEnv := os.Getenv("SYSLOGRELAYD_SYSLOG_PORT")

	ports := strings.Split(syslogPortEnv, ",")

	for _, port := range ports {
		if convertedPort, err := strconv.ParseInt(port, 10, 16); err == nil {
			syslogPorts = append(syslogPorts, uint16(convertedPort))
		}
	}

	if len(syslogPorts) == 0 {
		syslogPorts = append(syslogPorts, syslogDefaultPort)
	}

	appConfig.syslogServer.SyslogServer = &net.UDPAddr{
		IP:   syslogIP,
		Port: int(syslogPorts[0]),
	}
	appConfig.syslogServer.SyslogServerPorts = syslogPorts
}
