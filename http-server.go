package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

var counter uint64

func main() {

	log.Print("http-server listening...")

	http.HandleFunc("/", root)
	http.HandleFunc("/metrics", metrics)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func root(w http.ResponseWriter, _ *http.Request) {
	atomic.AddUint64(&counter, 1)
	w.Write([]byte("World"))
}

func metrics(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w,
		"%s\n%s\n%s %d",
		"# HELP cpa_request_total request total",
		"# TYPE cpa_request_total counter",
		"http_request_total{name=\"controller_metrics\",} ",
		counter)
}
