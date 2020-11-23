package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

// hey -n 400000 -c 50 http://localhost:8081/
// wrk -t3 -c50 -d10s --latency http://127.0.0.1:8081/
func main() {
	log.Print("http-server listening...")
	log.Fatal(fasthttp.ListenAndServe(":8081", requestHandler))
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("World")
}
