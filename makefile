PROJECTNAME = go-ssh

VERSION = 1.0.0


# install
install : linux windows darwin

## linux: 编译打包linux
.PHONY: linux
linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/$(PROJECTNAME)$(VERSION).linux-amd64

## windows: 编译打包windows
.PHONY: windows
windows:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/$(PROJECTNAME)$(VERSION).windows-amd64.exe

## mac: 编译打包mac
.PHONY: darwin
darwin:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/$(PROJECTNAME)$(VERSION).darwin-amd64

clean:
	rm -rf ./bin/*