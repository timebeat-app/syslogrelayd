package http_server

import (
	"encoding/json"
	"fmt"
	"github.com/timebeat-app/syslogrelayd/syslog_client"
	"io"
	"log"
	"net/http"
	"strings"
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

		body, err := io.ReadAll(httpRequest.Body)
		if err != nil {
			fmt.Printf("Error reading http request: %s\n", err)
			return
		}

		alertMessage := parseJson(body)
		for _, alert := range alertMessage.Alerts {
			syslogMessage := formatSyslogMessage(alert)
			controller.syslogController.Log(syslogMessage)
		}

		return
	} else {
		fmt.Printf("Can only accept accept POST. We got: %s\n", httpRequest.Method)
	}
}

func formatSyslogMessage(alert Alert) string {

	var msgbuilder strings.Builder

	msgbuilder.WriteString(fmt.Sprintf("Timebeat Alert - status: %s, ", strings.ToUpper(alert.Status)))

	if hostname, ok := alert.Labels["host.name"]; ok {
		msgbuilder.WriteString(fmt.Sprintf("verification server: %s, ", strings.ToLower(hostname)))
	}

	if peerId, ok := alert.Labels["clock_sync.source.peer_identity.id"]; ok {
		msgbuilder.WriteString(fmt.Sprintf("ptp peer id: %s, ", strings.ToLower(peerId)))
	}

	if alertName, ok := alert.Labels["alertName"]; ok {

		msgbuilder.WriteString(fmt.Sprintf("alert: %s, ", strings.ToLower(alertName)))
	}

	if !alert.StartsAt.IsZero() {
		msgbuilder.WriteString(fmt.Sprintf("starts at: %s, ", alert.StartsAt.String()))
	}

	if !alert.EndsAt.IsZero() {
		msgbuilder.WriteString(fmt.Sprintf("ends at: %s, ", alert.EndsAt.String()))
	}

	msgbuilder.WriteString(fmt.Sprintf("url: %s", alert.GeneratorURL))
	return msgbuilder.String()
}

func handleReadinessProbe(responseWriter http.ResponseWriter, _ *http.Request) {
	responseWriter.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(responseWriter, "syslogrelayd is healthy")
}

func parseJson(jsonData []byte) *AlertMessage {

	var alertMessage AlertMessage
	err := json.Unmarshal(jsonData, &alertMessage)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}

	return &alertMessage
}
