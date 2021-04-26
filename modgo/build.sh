#!/bin/sh
set -xe

# 在本地生成生产环境可执行文件

if [[ "pro" == $1 ]]; then
  #  生产环境
  tag=$(date +%Y%m%d_%H%M_pro_)$(git rev-parse --short HEAD)
else
  #  测试环境
  tag=$(date +%Y%m%d_%H%M_test_)$(git rev-parse --short HEAD)
fi

GOOS=linux GOARCH=amd64 go build -o main tmain.go

docker build --no-cache -t gotest:$tag .

rm -f main