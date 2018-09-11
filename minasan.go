package main

import (
	"fmt"
	"os"
	"os/signal"

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
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for sig := range c {
		// sig is a ^C, handle it
		fmt.Println(sig)
		return
	}
}
