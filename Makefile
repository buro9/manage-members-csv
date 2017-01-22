.PHONY: native
native:
	go build -v && file ./manage-members-csv

.PHONY: linux
linux:
	env GOOS=linux GOARCH=amd64 go build -v && file ./manage-members-csv

.PHONY: osx
osx:
	env GOOS=darwin GOARCH=amd64 go build -v && file ./manage-members-csv

.PHONY: windows
windows:
	env GOOS=windows GOARCH=amd64 go build -v && file ./manage-members-csv
