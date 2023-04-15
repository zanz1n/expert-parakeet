build:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o dist/darwin-amd64 .
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/linux-amd64 .
	GOOS=freebsd GOARCH=amd64 go build -ldflags "-s -w" -o dist/freebsd-amd64 .
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o dist/windows-amd64.exe .
