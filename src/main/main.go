package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func main() {
	port := ":8080"

	templates := populateTemplates()
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			requestedFile := r.URL.Path[1:]
			template := templates.Lookup(requestedFile + ".html")
			fmt.Printf(requestedFile)

			if template != nil {
				template.Execute(w, nil)
			} else {
				w.WriteHeader(404)
			}

		})

	fmt.Printf("Serving on port " + port)
	http.ListenAndServe(port, nil)

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
