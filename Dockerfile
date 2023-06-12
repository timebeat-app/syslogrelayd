FROM golang:1.20.5-alpine
LABEL authors="Lasse @ Timebeat.app"
WORKDIR /app
COPY src/go.mod ./src/
COPY src/http_server ./src/
COPY src/syslog_client ./src/
COPY src/*.go ./src/
WORKDIR /app/src
ENV GOPATH=/app/src
RUN go build -o /syslogrelayd
EXPOSE 8080
ENTRYPOINT [ "/syslogrelayd" ]