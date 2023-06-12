FROM golang:1.20.5-alpine
LABEL authors="Lasse @ Timebeat.app"
WORKDIR /app
COPY go.mod ./
COPY http_server /usr/local/go/src/
COPY syslog_client /usr/local/go/src/
COPY *.go ./
RUN go build -o /syslogrelayd
EXPOSE 8080
ENTRYPOINT [ "/syslogrelayd" ]