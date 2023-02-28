package main

import (
	"errors"
	"log"
	"testing"
	"time"

	"github.com/getsentry/sentry-go"
)

func TestSentry(t *testing.T) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://aaaa",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)
	err = errors.New("aaaaaaaaaIt works!")
	sentry.CaptureException(err)
	sentry.CaptureMessage("It works!")
}
