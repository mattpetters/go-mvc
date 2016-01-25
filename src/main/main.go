package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"
	"viewmodels"
)

func main() {
	port := ":8080"

	templates := populateTemplates()
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			requestedFile := r.URL.Path[1:]
			template := templates.Lookup(requestedFile + ".html")
			fmt.Printf(requestedFile)

			var context interface{}
			if requestedFile == "index" {
				context = viewmodels.GetHome()
			}
			if template != nil {
				template.Execute(w, context)
			} else {
				w.WriteHeader(404)
			}

		})

	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)

	fmt.Printf("Serving on port " + port)
	http.ListenAndServe(port, nil)

}

func serveResource(w http.ResponseWriter, r *http.Request) {
	path := "public" + r.URL.Path
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err == nil {
		//have to free up resources whenever opening files
		defer f.Close()

		w.Header().Add("Content Type", contentType)

		bufferedReader := bufio.NewReader(f)
		bufferedReader.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}

func populateTemplates() *template.Template {
	result := template.New("templates")

	basePath := "templates"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)

	templatePaths := new([]string)

	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		}
	}

	result.ParseFiles(*templatePaths...)
	return result
}
