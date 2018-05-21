# This how we want to name the binary output
BINARY=test.exe
# These are the values we want to pass for VERSION  and BUILD
VERSION=`git describe --tags`
TIME=`date +%FT%T%z`
MODE=$(mode)
ifeq ($(MODE),)
	MODE=`debug`
endif
# Setup the -Idflags options for go build here,interpolate the variable values
LDFLAGS=-ldflags "-X main.BuildVersion=${VERSION} -X main.BuildTime=${TIME} -X main.BuildMode=${MODE}"
# Builds the project
build:
	go build ${LDFLAGS} -o ${BINARY}

clean:
	rm -rf ${BINARY}
	rm -rf dist

deps:
	go get -u -v github.com/gin-gonic/gin
	go get -u -v github.com/go-ini/ini

dist:
	mkdir dist
	cp ${BINARY} dist/
	mkdir dist/configs
	cp configs/config_${MODE}_.ini dist/configs/

.PHONY:  build clean dist