package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./public")))
	http.ListenAndServe(":8000", mux)
}
func foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Foo"))
}
