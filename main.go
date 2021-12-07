package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8000", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	Hostnam, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "<!DOCTYPE html><html><head><style>table {font-family: arial;border-collapse: collapse;width: 100%;font-size: 10px}td, th {border: 2px solid #FFFFE0;text-align: left;padding: 4px;}tr:nth-child(even) {background-color: #20B2AA;}tr:nth-child(odd) {background-color: #6495ED;}</style></head><body>")
	fmt.Fprintf(w, "<h1>YESSERV v2</h1>")
	//proceed with HTTP parameters
	fmt.Fprintf(w, "<table><tr><th>HTTP parameter</th><th>Value</th></tr>")
	fmt.Fprintf(w, "<tr><th>Container Instance name</th><th>%q</th></tr>", []byte(Hostnam))
	fmt.Fprintf(w, "<tr><th>Host Header</th><th>%q</th></tr>", r.Host)
	fmt.Fprintf(w, "<tr><th>Remote Address</th><th>%q</th></tr>", r.RemoteAddr)
	fmt.Fprintf(w, "<tr><th>Protocol version</th><th>%s</th></tr>", r.Proto)
	fmt.Fprintf(w, "<tr><th>Method</th><th>%s</th></tr>", r.Method)
	fmt.Fprintf(w, "<tr><th>URI</th><th>%s</th></tr>", r.URL)
	fmt.Fprintf(w, "<tr><th>Request(original) URI</th><th>%s</th></tr>", r.RequestURI)
	fmt.Fprintf(w, "<tr><th>Trailer</th><th>%s</th></tr>", r.Trailer)
	fmt.Fprintf(w, "<tr><th>ContentLength</th><th>%d</th></tr>", r.ContentLength)
	fmt.Fprintf(w, "<tr><th>TransferEncoding</th><th>%s</th></tr>", r.TransferEncoding)
	fmt.Fprintf(w, "<tr><th>Close</th><th>%v</th></tr>", r.Close)
	fmt.Fprintf(w, "<tr><th>Form Data</th><th>%s</th></tr>", r.Form)
	fmt.Fprintf(w, "</table><hr><table>")
	//Iterate over all HTTP header fields
	for k, v := range r.Header {
		fmt.Fprintf(w, "<tr><th>%q</th><th>%q</th></tr>", k, v)
	}
	fmt.Fprintf(w, "</table><hr><table>")
	fmt.Fprintf(w, "<tr><th>Body</th><th>%s</th></tr>", r.Body)
	fmt.Fprintf(w, "</table></body></html>")
}
