SHELL := /bin/bash
cwd  := $(shell pwd)
outfile := coverage

test:
	bash ./scripts/test.sh --open

docs:
	godoc -http=:8080