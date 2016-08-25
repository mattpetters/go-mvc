package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
	"controllers"
)

func main() {
	port := ":8000"

	templates := populateTemplates()
	controllers.Register(templates)

	fmt.Printf("Serving on port " + port + "\n")
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
