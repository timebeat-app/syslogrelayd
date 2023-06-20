package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting syslogrelayd")
	appController := NewAppController()
	appController.Run()
}
