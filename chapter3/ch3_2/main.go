package main

import (
	"go_programming_language/chapter3/ch3_2/surface"
	"net/http"
)

func main() {
	surface.Surface()
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "image/svg+xml")
		writer.Write([]byte(surface.Surface()))
	})
	http.ListenAndServe("localhost:8080", nil)
}
