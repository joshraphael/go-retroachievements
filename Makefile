SHELL := /bin/bash
cwd  := $(shell pwd)
outfile := coverage

test:
	bash ./scripts/test.sh --open

smoke:
	bash ./scripts/smoke.sh

docs:
	godoc -http=:8080