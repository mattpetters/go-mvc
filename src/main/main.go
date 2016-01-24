package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := ":8080"
	http.Handle("/", new(HTTPHandler))
	fmt.Printf("Go is serving on port" + port + "...")
	http.ListenAndServe(port, nil)

}

type HTTPHandler struct {
	http.Handler
}

func (this *HTTPHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	path := "public" + request.URL.Path

	//data, err := ioutil.ReadFile(string(path))

	f, err := os.Open(path)

	fmt.Printf(path)

	if err == nil {
		bufferedReader := bufio.NewReader(f)

		var contentType string

		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else {
			contentType = "text/plain"
		}

		writer.Header().Add("Content Type", contentType)
		//writer.Write(data)
		bufferedReader.WriteTo(writer)
	} else {
		writer.WriteHeader(404)
		writer.Write([]byte("404 - " + http.StatusText(404)))
	}
}
