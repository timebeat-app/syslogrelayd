package http_server

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

}
