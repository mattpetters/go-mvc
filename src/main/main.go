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
			context := Context{"this is the message"}
			tmpl.Execute(w, context)
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
    <h1>This is an HTML page being served by Go</h1>
		<h2>{{.Message}}</h2>
    <img src="http://pngimg.com/upload/dog_PNG2422.png" alt="" />
  </body>
	</html>
`

type Context struct {
	Message string
}
