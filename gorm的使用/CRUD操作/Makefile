#!/usr/bin/env bash

test:fmt clean build

dev:fmt clean build run fmt-output

all:fmt build run fmt-output

fmt:
	goimports -l -w ./

build:
	go build -o output/curl ./

run:
	output/curl

clean:
	rm -rf output/curl
