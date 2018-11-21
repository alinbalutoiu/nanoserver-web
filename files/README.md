## Building the binaries

Windows:

``
GOOS=windows GOARCH=amd64 go build SimpleHTTPServer.go
``

Linux:

``
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o SimpleHTTPServer SimpleHTTPServer.go
``
