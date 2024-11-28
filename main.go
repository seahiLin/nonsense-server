package main

import (
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"buf.build/gen/go/nonsense/nonsense-api/connectrpc/go/agent/services/servicesconnect"
)

type AgentServer struct{}

func main() {
	server := &AgentServer{}
	mux := http.NewServeMux()
	path, handler := servicesconnect.NewAgentServiceHandler(server)
	mux.Handle(path, handler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatal(err)
	}
}
