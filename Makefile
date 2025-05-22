# Makefile
app_name := uala-followers-service
version ?= latest

.PHONY: build

build:
	docker build -t $(app_name):$(version) .