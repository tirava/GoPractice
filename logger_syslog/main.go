package main

import (
	"log"
	"log/syslog"
)

func main() {
	logger, err := syslog.New(syslog.LOG_LOCAL3, "klimka") // may be Dial
	if err != nil {
		log.Fatal("Cannot attach to syslog")
	}
	defer logger.Close()

	logger.Debug("Debug Klim.")
	logger.Notice("Notice Klim.")
	logger.Warning("Warning Klim.")
	logger.Alert("Alert Klim.")
}
