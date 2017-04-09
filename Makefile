.PHONY: build doc fmt lint test vendor_clean vendor_get vendor_update

# This how we want to name the binary output
BINARY=msu-go-11-game

# These are the values we want to pass for VERSION and BUILD
VERSION=1.0.0
BUILD=`git rev-parse HEAD`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

# Prepend our _vendor directory to the system GOPATH
# so that import path resolution will prioritize
# our third party snapshots.
GOPATH := ${PWD}/_vendor:${GOPATH}
export GOPATH

# Default target
.DEFAULT_GOAL: $(BINARY)

# Builds the project
$(BINARY):
	go build ${LDFLAGS} -v -o ./bin/${BINARY} ./src/${BINARY}.go

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	golint ./src

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	go fmt ./src/...

test:
	go test ./src/...

# Installs our project: copies binaries
install:
	go install ${LDFLAGS} -o ./bin/${BINARY}

# Cleans our project: deletes binaries
clean:
	if [ -f ./bin/${BINARY} ] ; then rm ./bin/${BINARY} ; fi

doc:
	godoc -http=:6060 -index

vendor_clean:
	rm -dRf ./_vendor/src

# We have to set GOPATH to just the _vendor
# directory to ensure that `go get` doesn't
# update packages in our primary GOPATH instead.
# This will happen if you already have the package
# installed in GOPATH since `go get` will use
# that existing location as the destination.
vendor_get: vendor_clean
	GOPATH=${PWD}/_vendor go get -d -u -v \
	github.com/jpoehls/gophermail \
	github.com/codegangsta/martini

vendor_update: vendor_get
	rm -rf `find ./_vendor/src -type d -name .git` \
	&& rm -rf `find ./_vendor/src -type d -name .hg` \
	&& rm -rf `find ./_vendor/src -type d -name .bzr` \
	&& rm -rf `find ./_vendor/src -type d -name .svn`
