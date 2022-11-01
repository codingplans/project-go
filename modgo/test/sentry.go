package main

import (
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

func TestSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://aaaa",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("It works!")
}
