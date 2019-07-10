package main

import (
	"os"

	"github.com/evalphobia/logrus_sentry"
	"github.com/factorysh/minasan/cmd"
	"github.com/factorysh/minasan/version"
	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
)

func main() {
	// logrus hook for sentry, if DSN is provided
	dsn := os.Getenv("SENTRY_DSN")
	if dsn != "" {
		// sentryHook := sentry.NewHook(dsn, log.PanicLevel, log.FatalLevel, log.ErrorLevel)
		sentryHook, err := logrus_sentry.NewWithTagsSentryHook(dsn, map[string]string{
			"version": version.Version(),
			"program": "Minasan",
		}, []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
		})
		if err != nil {
			log.Error("sentry hook failed")
		}
		log.AddHook(sentryHook)
	}

	// Logrus hook for adding file name and line to logs
	filenameHook := filename.NewHook()
	log.AddHook(filenameHook)
	log.SetLevel(log.InfoLevel)
	log.Error("TEST ERROR")
	cmd.Execute()
}
