SHELL:=/bin/bash
.ONESHELL:

# go test -v ./...
test:
	set -euo pipefail
	go clean -testcache && \
	go test -v . | sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''
.PHONY:test

format:
	go fmt .

run:
	cat deluge-console-v-d.txt | go run main.go

NAME:=dcJSON
build-all:
	mkdir -p build ; \
	for os in darwin linux windows; do \
	for arch in amd64 arm64; do \
	output="build/$(NAME)-$$os-$$arch" ; \
	if [ "$${os}" == "windows" ]; then output="$${output}.exe"; fi ; \
	echo "building: $$output" ; \
	GOOS=$$os GOARCH=$$arch go build -o "$${output}" main.go ; \
	done ; \
	done