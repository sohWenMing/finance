package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sohWenMing/finance/internal/server"
)

func main() {
	doneChan := make(chan struct{})
	serverExitChan := make(chan struct{})
	exitChan := make(chan struct{})
	portChan := make(chan int)

	go server.InitServer(false, doneChan, serverExitChan, portChan)
	evalPort := <-portChan
	fmt.Println("server started: listening on port: ", evalPort)
	fmt.Println("enter exit to exit program")
	go initOsStdOutExit(exitChan)
	<-exitChan
	doneChan <- struct{}{}
	<-serverExitChan
	os.Exit(1)
}

func initOsStdOutExit(exitChan chan<- struct{}) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if strings.TrimSpace(input) == "exit" {
			exitChan <- struct{}{}
			return
		}
	}
}
