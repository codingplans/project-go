#!/bin/sh

# 使用方法： ./local_run.sh caishenweather
# $1 caishenweather

cd ~/go/src/$1

swag init

cd ~/go/src/testgo/modgo/swag
cp ~/go/src/$1/docs/docs.go ./docs

go run main.go
