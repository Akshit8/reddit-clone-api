// Package logger defines impl of app logger
package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// ConfigureAppLogger setups the logger
func ConfigureAppLogger() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}