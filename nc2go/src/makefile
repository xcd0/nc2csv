
BIN=../build/nc2go.exe

all: build

build:
	go build -o $(BIN)

run: build
	$(BIN) ./test/nc10

.PHONY: test
test: build
	go test

update:
	`cd ../../ > /dev/null 2>&1 && make > /dev/null 2>&1 && cd scanner/src > /dev/null 2>&1` > /dev/null 2>&1 &

