package http_server

import "time"

type HttpServerConfig struct {
	HTTPServerPort int
	HTTPURLPath    string
}

type Alert struct {
	Status       string            `json:"status"`
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
	StartsAt     time.Time         `json:"startsAt"`
	EndsAt       time.Time         `json:"endsAt"`
	GeneratorURL string            `json:"generatorURL"`
	Fingerprint  string            `json:"fingerprint"`
	SilenceURL   string            `json:"silenceURL"`
	DashboardURL string            `json:"dashboardURL"`
	PanelURL     string            `json:"panelURL"`
	Values       map[string]int    `json:"values"`
	ValueString  string            `json:"valueString"`
}

type AlertMessage struct {
	Receiver          string            `json:"receiver"`
	Status            string            `json:"status"`
	Alerts            []Alert           `json:"alerts"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	TruncatedAlerts   int               `json:"truncatedAlerts"`
	OrgId             int               `json:"orgId"`
	Title             string            `json:"title"`
	State             string            `json:"state"`
	Message           string            `json:"message"`
}

/*
{
	"receiver": "Timebeat Alert",
	"status": "resolved",
	"alerts": [
		{
			"status": "resolved",
			"labels": {
				"Timebeat": "Timebeat",
				"alertname": "Offset above 50us ",
				"clock_sync.source.peer_identity.id": "ipcchi4-pgm001",
				"grafana_folder": "Alerts",
				"host.name": "tvs1-ipcchi4",
				"notification": "all"
			},
			"annotations": {},
			"startsAt": "2024-10-31T14:05:00Z",
			"endsAt": "2024-10-31T14:07:40.00775909Z",
			"generatorURL": "https://ipc.timebeat.app/alerting/grafana/evtSSnuSk/view",
			"fingerprint": "97a09e1046098984",
			"silenceURL": "https://ipc.timebeat.app/alerting/silence/new?alertmanager=grafana\u0026matcher=Timebeat%!D(MISSING)Timebeat\u0026matcher=alertname%!D(MISSING)Offset+above+50us+\u0026matcher=clock_sync.source.peer_identity.id%!D(MISSING)ipcchi4-pgm001\u0026matcher=grafana_folder%!D(MISSING)Alerts\u0026matcher=host.name%!D(MISSING)tvs1-ipcchi4\u0026matcher=notification%!D(MISSING)all",
			"dashboardURL": "",
			"panelURL": "",
			"values": {
				"B": 165,
				"C": 1
			},
			"valueString": "[ var='C' labels={clock_sync.source.peer_identity.id=ipcchi4-pgm001, host.name=tvs1-ipcchi4} value=1 ], [ var='B' labels={clock_sync.source.peer_identity.id=ipcchi4-pgm001, host.name=tvs1-ipcchi4} value=165 ]"
		}
	],
	"groupLabels": {},
	"commonLabels": {
		"Timebeat": "Timebeat",
		"alertname": "Offset above 50us ",
		"clock_sync.source.peer_identity.id": "ipcchi4-pgm001",
		"grafana_folder": "Alerts",
		"host.name": "tvs1-ipcchi4",
		"notification": "all"
	},
	"commonAnnotations": {},
	"externalURL": "https://ipc.timebeat.app/",
	"version": "1",
	"groupKey": "{}/{Timebeat=\"Timebeat\"}:{}",
	"truncatedAlerts": 8,
	"orgId": 1,
	"title": "[RESOLVED]  (Offset above 50us  Timebeat ipcchi4-pgm001 Alerts tvs1-ipcchi4 all)",
	"state": "ok",
	"message": "**Resolved**\n\nValue: B=165, C=1\nLabels:\n - alertname = Offset above 50us \n - Timebeat = Timebeat\n - clock_sync.source.peer_identity.id = ipcchi4-pgm001\n - grafana_folder = Alerts\n - host.name = tvs1-ipcchi4\n - notification = all\nAnnotations:\nSource: https://ipc.timebeat.app/alerting/grafana/evtSSnuSk/view\nSilence: https://ipc.timebeat.app/alerting/silence/new?alertmanager=grafana\u0026matcher=Timebeat%!D(MISSING)Timebeat\u0026matcher=alertname%!D(MISSING)Offset+above+50us+\u0026matcher=clock_sync.source.peer_identity.id%!D(MISSING)ipcchi4-pgm001\u0026matcher=grafana_folder%!D(MISSING)Alerts\u0026matcher=host.name%!D(MISSING)tvs1-ipcchi4\u0026matcher=notification%!D(MISSING)all\n"
}


*/
