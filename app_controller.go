package main

import (
	"github.com/timebeat-app/syslogrelayd/http_server"
	"github.com/timebeat-app/syslogrelayd/syslog_client"
	"net"
)

type AppController struct {
	done                   chan struct{}
	httpServerController   *http_server.Controller
	syslogClientController *syslog_client.Controller
}

func NewAppController() *AppController {

	done := make(chan struct{})

	syslogHost := &net.UDPAddr{
		IP:   net.ParseIP("10.101.101.251"),
		Port: 514,
	}

	appController := &AppController{
		done:                   done,
		syslogClientController: syslog_client.NewSyslogClient(done, syslogHost),
	}
	appController.httpServerController = http_server.NewHttpServer(done, appController.syslogClientController)

	return appController
}

func (appController *AppController) Run() {

	appController.httpServerController.Run()
}
