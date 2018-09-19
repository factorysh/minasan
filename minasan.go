package main

import (
	"os"

	"github.com/factorysh/minasan/cmd"
	"github.com/onrik/logrus/filename"
	"github.com/onrik/logrus/sentry"
	log "github.com/sirupsen/logrus"
)

func main() {
	// logrus hook for sentry, if DSN is provided
	dsn := os.Getenv("SENTRY_DSN")
	if dsn != "" {
		sentryHook := sentry.NewHook(dsn, log.PanicLevel, log.FatalLevel, log.ErrorLevel)
		log.AddHook(sentryHook)
	}
	// Logrus hook for adding file name and line to logs
	filenameHook := filename.NewHook()
	log.AddHook(filenameHook)

	log.SetLevel(log.InfoLevel)
	cmd.Execute()
}
