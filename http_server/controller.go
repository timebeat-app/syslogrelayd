package http_server

import (
	"fmt"
	"log"
	"net/http"
)

type Controller struct {
	done chan struct{}
}

func NewHttpServer(done chan struct{}) *Controller {

	controller := &Controller{
		done: done,
	}
	return controller
}

func (controller *Controller) Run() {

	http.HandleFunc("/", handleWebhook)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func handleWebhook(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "You are on the home page")
}
