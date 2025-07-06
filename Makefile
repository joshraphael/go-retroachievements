SHELL := /bin/bash
cwd  := $(shell pwd)
outfile := coverage

test:
	bash ./scripts/test.sh --open

smoke:
	bash ./scripts/smoke.sh

docs: docs-gen
	cd _site && python3 -m http.server 8080

docs-gen:
	rm -rf _site/
	doc2go -embed -out _site ./...