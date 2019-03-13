#!/usr/bin/env bash
t=$1

if [ -z $t ];then
    go test
else
    go test -bench=. -benchmem
fi