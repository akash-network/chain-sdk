package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"pkg.akt.dev/go/testutil/mock"
)

func main() {
	flag.Parse()

	server, err := mock.NewServer(mock.Config{})
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	fmt.Fprintf(os.Stdout, "gateway: %s\n", server.GatewayURL())
	fmt.Fprintf(os.Stdout, "grpc: %s\n", server.GRPCAddr())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	if err := server.Stop(); err != nil {
		log.Fatalf("Error stopping server: %v", err)
	}
}
