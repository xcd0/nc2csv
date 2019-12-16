

all: build

build:
	go build -o scanner.exe

run: build
	./scanner nc10

update:
	`cd ../../ && make && cd scanner/src` > /dev/null &
