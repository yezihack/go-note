#!/bin/bash

fmt:
	gofmt -l -w ./

push:
	git config user.name 'lesmoon'
	git config user.email 'freeit@126.com'
	git add . && git commit -m "Study Go" && git push origin master
