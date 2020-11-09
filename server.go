package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n\n", r.RemoteAddr)
	fmt.Fprintf(w, "Proto = %s\n", r.Proto)
	fmt.Fprintf(w, "Method = %s\n", r.Method)
	fmt.Fprintf(w, "URL = %s\n", r.URL)
	fmt.Fprintf(w, "RequestURI = %s\n\n", r.RequestURI)
	fmt.Fprintf(w, "Trailer = %s\n\n", r.Trailer)
	//Iterate over all header fields
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header: %q, Value: %q\n", k, v)
	}
	fmt.Fprintf(w, "\n\n")
	fmt.Fprintf(w, "ContentLength = %d\n", r.ContentLength)
	fmt.Fprintf(w, "TransferEncoding = %s\n", r.TransferEncoding)
	fmt.Fprintf(w, "Close = %v\n\n", r.Close)
	fmt.Fprintf(w, "Form = %s\n\n", r.Form)
	fmt.Fprintf(w, "\n\n")
	fmt.Fprintf(w, "Body = %s\n\n", r.Body)
}
