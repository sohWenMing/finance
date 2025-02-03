package server

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/sohWenMing/finance/internal/handlers"
)

func InitServer(isTest bool, doneChan <-chan struct{}, serverExitChan chan<- struct{}, portChan chan<- int) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", handlers.PingHandler)

	port := ":8080"
	if isTest {
		port = ":0"
	}

	server := &http.Server{
		Handler: mux,
	}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	evalPort := listener.Addr().(*net.TCPAddr).Port
	portChan <- evalPort
	go func() {
		serveErr := server.Serve(listener)
		if serveErr != nil && !errors.Is(serveErr, http.ErrServerClosed) {
			log.Fatal(serveErr)
		}
	}()
	<-doneChan
	fmt.Println("closing server")
	server.Close()
	serverExitChan <- struct{}{}
}
