package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestRunMain(t *testing.T) {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGUSR2, syscall.SIGUSR1)
	done := make(chan bool, 1)

	go func() {

		sig := <-sigs
		switch sig.Signal() {
		case syscall.SIGUSR1:
			break

		case syscall.SIGUSR2:
			break

		default:
			break
		}
		fmt.Println(sig)
		done <- true
	}()

	go main()
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

}
