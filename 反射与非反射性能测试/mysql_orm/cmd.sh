#!/usr/bin/env bash
# 1ns = 十亿分之一秒
go test -bench=. -cpu=1,2,3,4 -count=1 -benchmem