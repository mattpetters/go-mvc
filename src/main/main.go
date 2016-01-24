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
			//context := Context{"this is the message"}
			tmpl.Execute(w, r.URL.Path)
		}
	})

	fmt.Printf("Serving on port " + port)
	http.ListenAndServe(port, nil)
}

const doc = `
	<html>
  <head>
    <title>Hello</title>
  </head>
  <body>
    {{if eq . "/Google"}}
		 	<h1>Google made go</h1>
		{{else}}
			<h1>Hello {{.}}</h1>
		{{end}}
  </body>
	</html>
`

// type Context struct {
// 	Message string
// }
