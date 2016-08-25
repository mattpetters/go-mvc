package controllers

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func Register(templates *template.Template) {
	// http.HandleFunc("/",
	// 	func(w http.ResponseWriter, r *http.Request) {
	// 		requestedFile := r.URL.Path[1:]
	// 		template := templates.Lookup(requestedFile + ".html")
	// 		fmt.Printf(requestedFile + "\n")
	//
	// 		var context interface{}
	// 		switch requestedFile {
	// 			case "home":
	// 				context = viewmodels.GetHome()
	// 			case "404":
	// 				context = viewmodels.Get404()
	// 			default:
	// 				template = templates.Lookup("404" + ".html")
	// 				context = viewmodels.Get404()
	// 				w.WriteHeader(404)
	// 				fmt.Printf("Showing 404 page \n")
	// 		}
	// 		if template != nil {
	// 			template.Execute(w, context)
	// 		} else {
	//
	// 		}
	// 	})
	
	//Home route
	hc := new(homeController)
	hc.template = templates.Lookup("home.html")
	http.HandleFunc("/home", hc.get)

	//Lights route
	lc := new(lightsController)
	lc.template = templates.Lookup("lights.html")
	http.HandleFunc("/lights", lc.get)

	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)
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
