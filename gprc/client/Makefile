#!/usr/bin/env bash

all: dev run

fmt:
	goimports -l -w ./app/

install: fmt clean

clean:
	rm -rf output/

dev: install
	go build -o output/chat_client ./app

test: install
	go build  -o output/chat_client ./app

run:dev
	ETCD_ADDR=127.0.0.1:12379 output/chat_client