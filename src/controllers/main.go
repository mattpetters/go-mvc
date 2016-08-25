package controller

import (
	"net/http"
	"text/template"
	"viewmodels"
)

func Register(templates *template.Template) {
	http.HandleFunc("/",
		func(w http.ResponseWriter, req *http.Request) {
			requestedFile := req.URL.Path[1:]
			template :=
				templates.Lookup(requestedFile + ".html")

			var context interface{} = nil
			switch requestedFile {
			case "home":
				context = viewmodels.GetHome()
			case "categories":
				context = viewmodels.GetCategories()
			case "products":
				context = viewmodels.GetProducts()
			case "product":
				context = viewmodels.GetProduct()
			}
			if template != nil {
				template.Execture(w, context)
			} else {
				w.WriteHeader(404)
			}
		})
}
