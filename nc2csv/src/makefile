
BIN=nc2csv

DST=../build
GOARCH=amd64
FLAGS_WIN=-ldflags='-H windowsgui -w -s -extldflags "-static"' -a -tags netgo -installsuffix netgo
FLAGS=-ldflags='-w -s -extldflags "-static"' -a -tags netgo -installsuffix netgo

all: build

build:
.PHONY: build
build:
	mkdir -p $(DST)
ifeq ($(shell uname -o),Msys)
	go build -o $(DST)/$(BIN).exe
else
	go build -o $(DST)/$(BIN)
endif
	cp -rf $(DST)/$(BIN) .

run: build
	rm -rf ./test/A12.ncd.csv
	$(DST)/$(BIN) ./test/A12.ncd
	cat ./test/A12.ncd.csv

.PHONY: test
test: build
	go test

update:
	`cd ../../ > /dev/null 2>&1 && make > /dev/null 2>&1 && cd scanner/src > /dev/null 2>&1` > /dev/null 2>&1 &

debug-build:
	GOARCH=$(GOARCH) GOOS=windows go build -o $(DST)/$(BIN)_windows.exe -ldflags="-H windowsgui"
	GOARCH=$(GOARCH) GOOS=darwin go build -o $(DST)/$(BIN)_macOS
	GOARCH=$(GOARCH) GOOS=linux go build -o $(DST)/$(BIN)_linux
	GOARCH=$(GOARCH) GOOS=freebsd go build -o $(DST)/$(BIN)_freeBSD

release-build:
	GOARCH=$(GOARCH) GOOS=windows go build -o $(DST)/$(BIN)_windows.exe $(FLAGS_WIN)
	GOARCH=$(GOARCH) GOOS=darwin go build -o $(DST)/$(BIN)_macOS $(FLAGS)
	GOARCH=$(GOARCH) GOOS=linux go build -o $(DST)/$(BIN)_linux $(FLAGS)
	GOARCH=$(GOARCH) GOOS=freebsd go build -o $(DST)/$(BIN)_freeBSD $(FLAGS)

release: release-build
	# for windows
	cd $(DST) && \
	mv $(BIN)_windows.exe $(BIN).exe && \
	zip $(BIN)_binary_$(GOARCH)_windows.zip $(BIN).exe && \
	mv $(BIN).exe $(BIN)_windows.exe
	# for mac
	cd $(DST) && \
	mv $(BIN)_macOS $(BIN) && \
	zip $(BIN)_binary_$(GOARCH)_macOS.zip $(BIN) && \
	mv $(BIN) $(BIN)_macOS
	# for linux
	cd $(DST) && \
	mv $(BIN)_linux $(BIN) && \
	zip $(BIN)_binary_$(GOARCH)_linux.zip $(BIN) && \
	mv $(BIN) $(BIN)_linux
	# for freeBSD
	cd $(DST) && \
	mv $(BIN)_freeBSD $(BIN) && \
	zip $(BIN)_binary_$(GOARCH)_freeBSD.zip $(BIN) && \
	mv $(BIN) $(BIN)_freeBSD
	# clean
	cd $(DST) && \
	rm $(DST)/$(BIN)_windows.exe && \
	rm $(DST)/$(BIN)_macOS && \
	rm $(DST)/$(BIN)_linux && \
	rm $(DST)/$(BIN)_freeBSD
