#!/usr/bin/env sh
cpu=$1
go test -bench=. -benchmem -cpu=$cpu