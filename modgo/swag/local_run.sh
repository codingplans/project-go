#!/bin/sh

# 使用方法： ./local_run.sh caishenweather
# $1 caishenweather

cd ~/go/src/$1
flag=0
# 这里的-f参数判断$myFile是否存在
if [ ! -f "~/go/src/$1/main.go" ]; then
cp cmd/main.go ./
flag=1
fi

swag init

cd ~/go/src/testgo/modgo/swag
cp ~/go/src/$1/docs/docs.go ./docs

if [ $flag==1 ];then
  rm ~/go/src/$1/main.go
fi

go run main.go

