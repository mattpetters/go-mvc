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
			tmpl.Execute(w, nil)
		}
	})

	fmt.Printf("Serving on port " + port)
	http.ListenAndServe(port, nil)
}

const doc = `
	<html>
  <head>
    <title>Hello World</title>
  </head>
  <body>
    <h1>This is an HTML page being served by Go</h1>
    <img src="http://pngimg.com/upload/dog_PNG2422.png" alt="" />
  </body>
	</html>
`
