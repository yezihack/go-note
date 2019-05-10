#!/usr/bin/env bash

all: dev run

fmt:
	goimports -l -w ./app/

install: fmt clean

clean:
	rm -rf output/

dev: install
	go build -o output/grpc_chat ./app

test: install
	go build  -o output/bin/go_love_coin ./app

run:
	ETCD_ADDR=127.0.0.1:12379 output/bin/go_love_coin

