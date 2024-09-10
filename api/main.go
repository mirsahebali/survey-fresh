package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./api/static")))
	http.ListenAndServe(":8000", mux)
}
