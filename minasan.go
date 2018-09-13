package main

import (
	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
)

import "gitlab.bearstech.com/factory/minasan/cmd"

func main() {
	// Logrus hook for adding file name and line to logs
	filenameHook := filename.NewHook()
	log.SetLevel(log.InfoLevel)
	log.AddHook(filenameHook)
	cmd.Execute()
}
