package main

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"buf.build/gen/go/nonsense/nonsense-api/connectrpc/go/agent/services/servicesconnect"
)

type AgentServer struct{}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := &AgentServer{}
	mux := http.NewServeMux()
	path, handler := servicesconnect.NewAgentServiceHandler(server)
	mux.Handle(path, handler)

	log.Println("Starting server on :" + port)
	if err := http.ListenAndServe(
		"0.0.0.0:"+port,
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatal(err)
	}
}
