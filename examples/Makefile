.DEFAULT_GOAL := overview

overview:
	@echo "A series of examples for different tools and techniques in this project\n"
	@echo "make layered   - run basic layered DDD example"
	@echo "make gotest    - run go test example"
.PHONY: overview

basic:
	go run layered/main.go
.PHONY: basic

gotest:
	go test -v ./gotest/...
.PHONY: gotest