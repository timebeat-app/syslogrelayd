package syslog_client

type Controller struct {
	done chan struct{}
}

func NewSyslogClient(done chan struct{}) *Controller {

	controller := &Controller{
		done: done,
	}
	return controller
}

func (controller *Controller) Run() {

}
