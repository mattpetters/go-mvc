package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":8080"
	fmt.Printf("Go is serving on port" + port + "...")
	http.ListenAndServe(port, http.FileServer(http.Dir("public")))

}
