package main

import (
	"context"
	"fmt"
	"github.com/akbarpambudi/go-point-of-sales/internal/library"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	server, closeWebService, err := library.NewWebService(ctx)
	if err != nil {
		log.Fatalf("something went wrong: %v", err)
	}

	defer closeWebService()

	osSignalCh := make(chan os.Signal)
	signal.Notify(osSignalCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-osSignalCh
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		closeWebService()
		cancel()
		os.Exit(0)
	}()

	listeningErr := http.ListenAndServe(":8080", server)
	if listeningErr != nil {
		log.Fatalf("something went wrong: %v", err)
	}
}
