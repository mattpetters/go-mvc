package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	port := ":8080"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content Type", "text/html")
		tmpl, err := template.New("test").Parse(doc)

		if err == nil {
			context := Context{
				[3]string{"Apple", "Orange", "Banana"},
				"a title",
			}
			tmpl.Execute(w, context)
		}
	})

	fmt.Printf("Serving on port " + port)
	http.ListenAndServe(port, nil)
}

const doc = `
	<html>
  <head>
    <title>{{.Title}}</title>
  </head>
  <body>
    <h1>List of fruit</h1>
		<ul>
			{{range .Fruit}}
				<li>{{.}}</li>
			{{end}}
		</ul>
  </body>
	</html>
`

type Context struct {
	Fruit [3]string
	Title string
}
