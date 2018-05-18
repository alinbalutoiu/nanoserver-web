package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(writer http.ResponseWriter, req *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		name = "HOSTNAME_UNKNOWN"
	}
	fmt.Fprintf(writer, "<h1>This request was processed by container: %s</h1>\n", name)
}

func main() {
	fmt.Fprintf(os.Stdout, "Web Server started. Listening on 0.0.0.0:80\n")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
