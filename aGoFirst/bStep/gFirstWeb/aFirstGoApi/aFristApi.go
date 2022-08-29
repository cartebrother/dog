package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello world, dog !"))
	})

	http.ListenAndServe(":8233", nil)
}
