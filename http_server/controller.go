package http_server

import (
	"fmt"
	"github.com/timebeat-app/syslogrelayd/syslog_client"
	"log"
	"net/http"
)

type Controller struct {
	done             chan struct{}
	syslogController *syslog_client.Controller
}

func NewHttpServer(done chan struct{}, syslogController *syslog_client.Controller) *Controller {

	controller := &Controller{
		done:             done,
		syslogController: syslogController,
	}
	return controller
}

func (controller *Controller) Run() {

	http.HandleFunc("/", controller.handleWebhook)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func (controller *Controller) handleWebhook(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Alert submitted")
	controller.syslogController.Log()
}
