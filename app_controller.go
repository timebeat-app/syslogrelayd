package main

import (
	"github.com/timebeat-app/syslogrelayd/http_server"
	"github.com/timebeat-app/syslogrelayd/syslog_client"
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
