UNAME := $(shell uname)
XARGS = xargs
ARCH ?= $(shell go env GOARCH)
INTERNAL_PROTO_FILES=$(shell find . -name *.proto)

# 这是一个测试 case 的实现
test:
	@echo $(INTERNAL_PROTO_FILES)
	mkdir www

.PHONY: config

# generate internal proto
config:

# 不存在时
ifneq ($(INTERNAL_PROTO_FILES), "")
	@echo 生成内部 $(INTERNAL_PROTO_FILES)
else
	@echo 不存在
endif

# 关于更多文档 查询 gnu 官网 https://www.gnu.org/software/make/manual/html_node/Conditional-Syntax.html
more:
	@echo 感谢使用

all:
	make test1


# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
