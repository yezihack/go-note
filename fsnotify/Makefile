#!/usr/bin/env bash

all: run

run:clean build
	./run/fsnotify

fmt:
	gofmt -l -w ./

build:fmt
	go build -v -o ./run/fsnotify ./

clean:
	rm -rf run/fsnotify
