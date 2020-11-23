package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

// hey -n 400000 -c 50 http://localhost:8081/
func main() {
	log.Print("http-server listening...")
	log.Fatal(fasthttp.ListenAndServe(":8081", requestHandler))
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("World")
}
