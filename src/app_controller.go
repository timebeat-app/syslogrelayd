package main

import (
	"syslogrelayd/http_server"
	"syslogrelayd/syslog_client"
)

type AppController struct {
	done                   chan struct{}
	httpServerController   *http_server.Controller
	syslogClientController *syslog_client.Controller
}

func NewAppController() *AppController {

	done := make(chan struct{})
	appController := &AppController{
		done:                   done,
		syslogClientController: syslog_client.NewSyslogClient(done),
		httpServerController:   http_server.NewHttpServer(done),
	}

	return appController
}

func (appController *AppController) Run() {

	appController.syslogClientController.Run()
}
