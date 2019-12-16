

all: build

build:
	go build -o scanner.exe

run: build
	./scanner nc10

update:
	`cd ../../ > /dev/null 2>&1 && make > /dev/null 2>&1 && cd scanner/src > /dev/null 2>&1` > /dev/null 2>&1 &

