package main

import (
	"fmt"

	"github.com/flashmob/go-guerrilla"
	"gitlab.bearstech.com/factory/minasan/processor"
)

func main() {
	d := guerrilla.Daemon{}
	d.AddProcessor("minasan", processor.MinasanProcessor)
	err := d.Start()

	if err == nil {
		fmt.Println("Server Started!")
	}
}
