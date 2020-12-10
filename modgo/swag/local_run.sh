#!/bin/sh

cd /Users/darren/go/src/calendarweather

swag init

cd /Users/darren/go/src/testgo/modgo/swag
cp /Users/darren/go/src/calendarweather/docs/docs.go ./docs

go run main.go
