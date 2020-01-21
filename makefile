
BIN=nc2csv

DST=./build
GOARCH=amd64

# 実行ファイルの大きさを小さくしたい場合コメントアウトを外す
#FLAGS_WIN=-ldflags='-H windowsgui -w -s -extldflags "-static"' -a -tags netgo -installsuffix netgo
#FLAGS=-ldflags='-w -s -extldflags "-static"' -a -tags netgo -installsuffix netgo

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
#	cp -rf $(DST)/$(BIN) .
#	cp -rf $(DST)/$(BIN) ..

release:
	rm -rf $(DST) && mkdir -p $(DST)
	# for windows
	GOARCH=$(GOARCH) GOOS=windows go build -o $(DST)/$(BIN)_windows.exe $(FLAGS_WIN)
	cd $(DST) && mv $(BIN)_windows.exe $(BIN).exe && zip $(BIN)_binary_$(GOARCH)_windows.zip $(BIN).exe && mv $(BIN).exe $(BIN)_windows.exe && rm $(BIN)_windows.exe
	# for mac
	GOARCH=$(GOARCH) GOOS=darwin go build -o $(DST)/$(BIN)_macOS $(FLAGS)
	cd $(DST) && mv $(BIN)_macOS $(BIN) && zip $(BIN)_binary_$(GOARCH)_macOS.zip $(BIN) && mv $(BIN) $(BIN)_macOS && rm $(BIN)_macOS
	# for linux
	GOARCH=$(GOARCH) GOOS=linux go build -o $(DST)/$(BIN)_linux $(FLAGS)
	cd $(DST) && mv $(BIN)_linux $(BIN) && zip $(BIN)_binary_$(GOARCH)_linux.zip $(BIN) && mv $(BIN) $(BIN)_linux && rm $(BIN)_linux
	# for freeBSD
	GOARCH=$(GOARCH) GOOS=freebsd go build -o $(DST)/$(BIN)_freeBSD $(FLAGS)
	cd $(DST) && mv $(BIN)_freeBSD $(BIN) && zip $(BIN)_binary_$(GOARCH)_freeBSD.zip $(BIN) && mv $(BIN) $(BIN)_freeBSD && rm $(BIN)_freeBSD

#run: build
#	rm -rf ./test/A12.ncd.csv
#	#$(DST)/$(BIN) ./test/A12.ncd
#	#cat ./test/A12.ncd.csv
#	$(DST)/$(BIN) ./test/nc10
#	cat ./test/nc10.csv

#.PHONY: test
#test: build
#	go test
#
#update:
#	`cd ../../ > /dev/null 2>&1 && make > /dev/null 2>&1 && cd scanner/src > /dev/null 2>&1` > /dev/null 2>&1 &

#debug-build:
#	GOARCH=$(GOARCH) GOOS=windows go build -o $(DST)/$(BIN)_windows.exe -ldflags="-H windowsgui"
#	GOARCH=$(GOARCH) GOOS=darwin go build -o $(DST)/$(BIN)_macOS
#	GOARCH=$(GOARCH) GOOS=linux go build -o $(DST)/$(BIN)_linux
#	GOARCH=$(GOARCH) GOOS=freebsd go build -o $(DST)/$(BIN)_freeBSD

#release: release-build
#	# for windows
#	cd $(DST) && \
#	mv $(BIN)_windows.exe $(BIN).exe && \
#	zip $(BIN)_binary_$(GOARCH)_windows.zip $(BIN).exe && \
#	mv $(BIN).exe $(BIN)_windows.exe
#	# for mac
#	cd $(DST) && \
#	mv $(BIN)_macOS $(BIN) && \
#	zip $(BIN)_binary_$(GOARCH)_macOS.zip $(BIN) && \
#	mv $(BIN) $(BIN)_macOS
#	# for linux
#	cd $(DST) && \
#	mv $(BIN)_linux $(BIN) && \
#	zip $(BIN)_binary_$(GOARCH)_linux.zip $(BIN) && \
#	mv $(BIN) $(BIN)_linux
#	# for freeBSD
#	cd $(DST) && \
#	mv $(BIN)_freeBSD $(BIN) && \
#	zip $(BIN)_binary_$(GOARCH)_freeBSD.zip $(BIN) && \
#	mv $(BIN) $(BIN)_freeBSD
#	# clean
#	cd $(DST) && \
#	rm $(DST)/$(BIN)_windows.exe && \
#	rm $(DST)/$(BIN)_macOS && \
#	rm $(DST)/$(BIN)_linux && \
#	rm $(DST)/$(BIN)_freeBSD
#	make build

