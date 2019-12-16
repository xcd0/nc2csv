

all: build

build:
	go build -o scanner.exe

run: build
	./scanner nc10

update:
	`cd ../../ > /dev/null && make > /dev/null && cd scanner/src > /dev/null` > /dev/null &

