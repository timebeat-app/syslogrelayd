FROM golang:1.20.5-alpine
LABEL authors="Lasse @ Timebeat.app"
WORKDIR /go/src/github.com/timebeat-app/syslogrelayd
COPY go.mod ./
COPY http_server ./http_server/
COPY syslog_client ./syslog_client/
COPY *.go ./
RUN go mod download
RUN go mod verify
RUN go build -o /syslogrelayd
EXPOSE 8080
ENTRYPOINT [ "/syslogrelayd" ]
