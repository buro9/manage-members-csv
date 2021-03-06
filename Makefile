
VERSION          := $(shell git describe --tags --always --dirty="-dev")
DATE             := $(shell date -u '+%Y-%m-%d-%H%M UTC')
VERSION_FLAGS    := -ldflags='-X "main.Version=$(VERSION)" -X "main.BuildTime=$(DATE)"'

.PHONY: native
native:
	go build -v $(VERSION_FLAGS) && file ./manage-members-csv

.PHONY: linux
linux:
	env GOOS=linux GOARCH=amd64 go build -v $(VERSION_FLAGS) -o manage-members-csv-linux && file ./manage-members-csv-linux

.PHONY: osx
osx:
	env GOOS=darwin GOARCH=amd64 go build -v $(VERSION_FLAGS) -o manage-members-csv-osx && file ./manage-members-csv-osx

.PHONY: windows
windows:
	env GOOS=windows GOARCH=amd64 go build -v $(VERSION_FLAGS) && file ./manage-members-csv.exe

.PHONY: all
all: native linux osx windows

.PHONY: clean
clean:
	rm -f manage-members-csv
	rm -f manage-members-csv-linux
	rm -f manage-members-csv-osx
	rm -f manage-members-csv.exe
