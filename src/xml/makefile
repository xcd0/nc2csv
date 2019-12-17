
BIN=read.exe

all: build

build:
	go build -o $(BIN)

run: build
	$(BIN) ./test/test.csv

.PHONY: test
test: build
	go test

update:
	`cd ../../ > /dev/null 2>&1 && make > /dev/null 2>&1 && cd scanner/src > /dev/null 2>&1` > /dev/null 2>&1 &

