package main

import (
	"context"
	"net/http"
	"net/http/pprof"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

func BenchmarkForSleep(b *testing.B) {
	for i := 0; i < b.N; {
		for {
			log.Info(time.Now())
			timer := time.NewTimer(time.Second)
			defer timer.Stop()
			select {
			case <-timer.C:
				go func() {
					defer func() {
						if r := recover(); r != nil {
							log.Info("Recovered in f", r)
						}
					}()
					m := make([]interface{}, 1000)
					log.Info(m)
				}()

			}

		}
	}
}

func BenchmarkForTime(b *testing.B) {
	for i := 0; i < b.N; {
		for {
			log.Info(time.Now())
			time.Sleep(time.Millisecond)
			go func() {
				defer func() {
					if r := recover(); r != nil {
						log.Info("Recovered in f", r)
					}
				}()
				m := make([]interface{}, 1000)
				log.Info(m)
			}()
		}
	}
}

func MakeDebugHandler(ctx context.Context) http.Handler {
	m := http.NewServeMux()
	m.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	m.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	m.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	m.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	m.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	return m
}

func ServDebug(addr string) {
	if addr == "" {
		log.Warnln("ServDebug need addr")
		return
	}
	ctx := context.TODO()
	r := MakeDebugHandler(ctx)
	e := http.ListenAndServe(addr, r)
	log.Warnf("ListenAndServe Error: %#v\n", e.Error())
}
