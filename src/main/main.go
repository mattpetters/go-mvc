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
		templates := template.New("template")
		templates.New("test").Parse(doc)
		templates.New("header").Parse(header)
		templates.New("footer").Parse(footer)
		context := Context{
			[3]string{"Apple", "Orange", "Banana"},
			"a title",
		}
		templates.Lookup("test").Execute(w, context)
	})

	fmt.Printf("Serving on port " + port)
	http.ListenAndServe(port, nil)
}

const doc = `
{{template "header" .Title}}
  <body>
    <h1>List of fruit</h1>
		<ul>
			{{range .Fruit}}
				<li>{{.}}</li>
			{{end}}
		</ul>
  </body>
	{{template "footer"}}
`
const header = `<html>
<head>
	<title>{{.}}</title>
</head>
`

const footer = `
	</html>
`

type Context struct {
	Fruit [3]string
	Title string
}
