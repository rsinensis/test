# This how we want to name the binary output
BINARY=test.exe
# These are the values we want to pass for VERSION  and BUILD
VERSION=`git describe --tags`
TIME=`date +%FT%T%z`
MODE=$(mode)
ifeq ($(MODE),)
	MODE=debug
endif
# Setup the -Idflags options for go build here,interpolate the variable values
LDFLAGS=-ldflags "-X main.BuildVersion=${VERSION} -X main.BuildTime=${TIME} -X main.BuildMode=${MODE}"
# Builds the project

.PHONY: build
build:
	@echo "build..."
	go build ${LDFLAGS} -o ${BINARY}

.PHONY: clean
clean:
	@echo "clean..."
	rm -rf ${BINARY}
	rm -rf debug
	rm -rf dist

.PHONY: deps
deps:
	@echo "deps..."
	go get -u -v github.com/gin-gonic/gin
	go get -u -v github.com/go-ini/ini
	go get -u -v github.com/go-sql-driver/mysql
	go get -u -v github.com/go-xorm/xorm
	go get -u -v github.com/gin-contrib/sessions
	go get -u -v github.com/swaggo/gin-swagger
	go get -u -v github.com/swaggo/gin-swagger/swaggerFiles

.PHONY: dist
dist: clean build
	@echo "dist..."
	mkdir dist
	cp ${BINARY} dist/
	mkdir dist/configs
	cp configs/config_"${MODE}".ini dist/configs/

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./.