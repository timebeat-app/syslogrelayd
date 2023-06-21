package http_server

import (
	"fmt"
	"github.com/timebeat-app/syslogrelayd/syslog_client"
	"io"
	"log"
	"net/http"
)

type Controller struct {
	done             chan struct{}
	syslogController *syslog_client.Controller
	httpServerConfig *HttpServerConfig
}

func NewHttpServer(done chan struct{}, httpServerConfig *HttpServerConfig,
	syslogController *syslog_client.Controller) *Controller {

	controller := &Controller{
		done:             done,
		syslogController: syslogController,
		httpServerConfig: httpServerConfig,
	}
	return controller
}

func (controller *Controller) Run() {

	http.HandleFunc("/health", handleReadinessProbe)
	http.HandleFunc(controller.httpServerConfig.HTTPURLPath, controller.handleWebhook)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", controller.httpServerConfig.HTTPServerPort), nil); err != nil {
		log.Fatal(err)
	}
}

func (controller *Controller) handleWebhook(_ http.ResponseWriter, httpRequest *http.Request) {

	if httpRequest.Method == http.MethodPost {
		if body, err := io.ReadAll(httpRequest.Body); err == nil {
			controller.syslogController.Log(body)
			return
		} else {
			fmt.Printf("Error reading http request: %s\n", err)
		}
	}
	fmt.Printf("Can only accept accept POST. We got: %s\n", httpRequest.Method)
}

func handleReadinessProbe(responseWriter http.ResponseWriter, _ *http.Request) {
	responseWriter.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(responseWriter, "syslogrelayd is healthy")
}
